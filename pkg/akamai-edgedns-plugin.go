/*
 * Copyright 2021 Akamai Technologies, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/grafana/grafana-plugin-sdk-go/backend"
	"github.com/grafana/grafana-plugin-sdk-go/backend/instancemgmt"
	"github.com/grafana/grafana-plugin-sdk-go/backend/log"
	"github.com/grafana/grafana-plugin-sdk-go/data"
)

type selectableValueStr struct {
	Label string `json:"label"`
	Value string `json:"value"`
}

// Label is displayed to the user in the query editor 'Reports' dropdown.
// Value is of the form 'path/metric', where 'path' is part of the OPEN API URL and 'metric' is the metric to graph.
var supportedReports = []selectableValueStr{
	selectableValueStr{Label: "Edge DNS traffic by time: hits", Value: "authoritative-dns-traffic-by-time/hits"},
	selectableValueStr{Label: "Edge DNS traffic by time: NXDOMAIN", Value: "authoritative-dns-traffic-by-time/nxdomain"},
}

// The datasource front-end sends zonenames (to graph) as a comma-separated string. OPEN API POST request needs a zonename list.
func zonesListFromZones(zoneNames string) []string {
	zoneNames = strings.Replace(zoneNames, " ", "", -1) // remove spaces
	rawList := strings.Split(zoneNames, ",")            // split on commas, may contain empty entries

	var cleanList []string
	for _, z := range rawList {
		if len(z) > 0 {
			cleanList = append(cleanList, z)
		}
	}
	return cleanList
}

// The datasource front-end sends the report to graph as a label/value pair.  The 'value' part has a format: path/metric',
// where 'path' is part of the OPEN API URL and 'metric' is the metric to graph.
// Given the (formatted) 'value', return the 'path' and 'metric' parts.
func pathAndMetricFromReport(report selectableValueStr) (string, string) {
	pm := strings.Split(report.Value, "/")
	return pm[0], pm[1] // reportPath, reportMetric
}

// The datasource configuration supplied by the front-end.
type dataSourceSettingsJson struct {
	ClientSecret string `json:"clientSecret"`
	Host         string `json:"host"`
	AccessToken  string `json:"accessToken"`
	ClientToken  string `json:"clientToken"`
}

// Query information supplied by the front-end
type dataQueryJson struct {
	DataSourceId   uint               `json:"dataSourceId"`
	IntervalMs     uint               `json:"intervalMs"`
	MaxDataPoints  uint               `json:"maxDataPoints"`
	SelectedReport selectableValueStr `json:"selectedReport"`
	ZoneNames      string             `json:"zoneNames"` // Comma-separated zone names
	MetricName     string             `json:"metricName"`
}

// Grafana structures and functions
func newDataSourceInstance(ctx context.Context, setting backend.DataSourceInstanceSettings) (instancemgmt.Instance, error) {
	return &AkamaiEdgeDnsDatasource{
		httpClient: &http.Client{},
	}, nil
}

type instanceSettings struct {
	httpClient *http.Client
}

// Called before creating a new instance to allow plugin to cleanup.
func (s *instanceSettings) Dispose() {
}

type AkamaiEdgeDnsDatasource struct {
	im         instancemgmt.InstanceManager
	httpClient *http.Client
}

func (td *AkamaiEdgeDnsDatasource) QueryData(ctx context.Context, req *backend.QueryDataRequest) (*backend.QueryDataResponse, error) {
	response := backend.NewQueryDataResponse()

	var dss dataSourceSettingsJson
	err := json.Unmarshal(req.PluginContext.DataSourceInstanceSettings.JSONData, &dss)
	if err != nil {
		return response, err
	}

	for _, q := range req.Queries {
		res, err := td.query(ctx, q, dss)
		if err != nil {
			// Create an error frame so the error shows in the panel
			errorFrame := data.NewFrame("Error")
			errorFrame.Meta = &data.FrameMeta{
				Type: data.FrameTypeTable,
			}
			errorFrame.Fields = append(errorFrame.Fields,
				data.NewField("Error", nil, []string{err.Error()}),
			)

			response.Responses[q.RefID] = backend.DataResponse{
				Frames: data.Frames{errorFrame},
				Error:  err,
			}
		} else {
			response.Responses[q.RefID] = *res
		}
	}

	return response, nil
}

func (td *AkamaiEdgeDnsDatasource) query(ctx context.Context, query backend.DataQuery, dss dataSourceSettingsJson) (*backend.DataResponse, error) {
	response := &backend.DataResponse{}

	var dqj dataQueryJson
	if err := json.Unmarshal(query.JSON, &dqj); err != nil {
		return nil, err
	}

	if len(dqj.SelectedReport.Label) == 0 || len(dqj.ZoneNames) == 0 {
		return nil, errors.New("Select a report and enter zone names")
	}

	interval := calculateInterval(query.TimeRange.From, query.TimeRange.To, dqj.MaxDataPoints)
	fromRounded, toRounded, err := adjustQueryTimes(query.TimeRange.From, query.TimeRange.To, interval)
	if err != nil {
		return nil, err
	}

	zoneNamesList := zonesListFromZones(dqj.ZoneNames)
	if len(zoneNamesList) == 0 {
		return nil, errors.New("Enter one or more zone names, separated by commas")
	}

	openApiRspDto, err := edgeDnsOpenApiQuery(zoneNamesList, fromRounded, toRounded, interval, dss.ClientSecret, dss.Host, dss.AccessToken, dss.ClientToken)
	if err != nil {
		return nil, err
	}

	numDataRows := len(openApiRspDto.Data)
	sampletime := make([]time.Time, numDataRows)
	hitspersec := make([]float64, numDataRows)

	_, reportMetric := pathAndMetricFromReport(dqj.SelectedReport)

	switch reportMetric {
	case "hits":
		for i, datum := range openApiRspDto.Data {
			sampletime[i], err = time.Parse(time.RFC3339, datum.StartDateTime)
			if err != nil {
				return nil, err
			}
			hitspersec[i], _ = strconv.ParseFloat(datum.SumHits, 64)
		}
	case "nxdomain":
		for i, datum := range openApiRspDto.Data {
			sampletime[i], err = time.Parse(time.RFC3339, datum.StartDateTime)
			if err != nil {
				return nil, err
			}
			hitspersec[i], _ = strconv.ParseFloat(datum.SumNxDomain, 64)
		}
	}

	frame := data.NewFrame("response")

	metricName := dqj.MetricName
	if len(metricName) == 0 {
		metricName = dqj.ZoneNames + " " + reportMetric
	}

	frame.Fields = append(frame.Fields, data.NewField("time", nil, sampletime))
	frame.Fields = append(frame.Fields, data.NewField(metricName, nil, hitspersec))

	response.Frames = append(response.Frames, frame)

	return response, nil
}

// The 'Save & Test' button on the datasource configuration page allows users to verify that the datasource is working as expected.
func (td *AkamaiEdgeDnsDatasource) CheckHealth(ctx context.Context, req *backend.CheckHealthRequest) (*backend.CheckHealthResult, error) {
	// log.DefaultLogger.Info("CheckHealth", "clientSecret", ds.ClientSecret)
	// log.DefaultLogger.Info("CheckHealth", "host", ds.Host)
	// log.DefaultLogger.Info("CheckHealth", "accessToken", ds.AccessToken)
	// log.DefaultLogger.Info("CheckHealth", "clientToken", ds.ClientToken)

	var ds dataSourceSettingsJson
	err := json.Unmarshal(req.PluginContext.DataSourceInstanceSettings.JSONData, &ds)
	if err != nil {
		return &backend.CheckHealthResult{
			Status:  backend.HealthStatusUnknown,
			Message: "Internal error. Failed to unmarshal healthcheck JSON",
		}, err
	}

	// Verify that the OPEN API responds.
	message, status := edgeDnsOpenApiHealthCheck(ds.ClientSecret, ds.Host, ds.AccessToken, ds.ClientToken)

	return &backend.CheckHealthResult{
		Status:  status,
		Message: message,
	}, nil

}

// 'CallResource' returns non-metric data to the front-end, in this case, the list of supported EdgeDns reports.
func (td *AkamaiEdgeDnsDatasource) CallResource(ctx context.Context, req *backend.CallResourceRequest, sender backend.CallResourceResponseSender) error {
	log.DefaultLogger.Info("CallResource", "path", req.Path)

	switch req.Path {
	case "datasource/resource/openapireports":
		b, err := json.Marshal(supportedReports)
		if err != nil {
			log.DefaultLogger.Error("Error marshalling reports", "err", err)
			return sender.Send(&backend.CallResourceResponse{
				Status: http.StatusInternalServerError,
				Body:   []byte(err.Error()),
			})
		}
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusOK,
			Body:   b,
		})
	default:
		return sender.Send(&backend.CallResourceResponse{
			Status: http.StatusNotFound,
			Body:   []byte("Unexpected resource URI: " + req.Path),
		})
	}
}
