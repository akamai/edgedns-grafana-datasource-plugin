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

import defaults from 'lodash/defaults';

import React, { ChangeEvent, PureComponent } from 'react';
import { QueryEditorProps, SelectableValue } from '@grafana/data';
import { InlineField, LegacyForms, AsyncSelect } from '@grafana/ui';
import { DataSource } from './DataSource';
import { defaultQuery, MyDataSourceOptions, MyQuery } from './types';

const { FormField } = LegacyForms;

type Props = QueryEditorProps<DataSource, MyQuery, MyDataSourceOptions>;

export class QueryEditor extends PureComponent<Props> {
  getReports = () => {
    const uri = 'datasource/resource/openapireports';
    return this.props.datasource.getResource(uri);
  };

  onSelectReportsChange = (item: SelectableValue<string>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, selectedReport: item });
    // console.log('selectedReport: ', item);
    if (item && query.zoneNames) {
      onRunQuery();
    }
  };

  onZoneNamesChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, zoneNames: event.target.value });
    // console.log('zoneNames: ' + event.target.value);
    if (event.target.value && query.selectedReport) {
      onRunQuery();
    }
  };

  onMetricNameChange = (event: ChangeEvent<HTMLInputElement>) => {
    const { onChange, query, onRunQuery } = this.props;
    onChange({ ...query, metricName: event.target.value });
    // console.log('metricName: ' + event.target.value);
    if (query.zoneNames && query.selectedReport) {
      onRunQuery();
    }
  };

  render() {
    const query = defaults(this.props.query, defaultQuery);
    const { selectedReport, zoneNames, metricName } = query;

    return (
      <div className="gf-form">
        <div>
          <InlineField label="Report">
            <AsyncSelect
              loadOptions={this.getReports}
              defaultOptions
              value={selectedReport}
              placeholder="Select a report"
              onChange={this.onSelectReportsChange}
            />
          </InlineField>
          <FormField
            value={zoneNames || ''}
            labelWidth={4}
            inputWidth={24}
            placeholder="Enter zone names"
            onChange={this.onZoneNamesChange}
            label="Zones"
            tooltip="Comma-separted zone names. Metrics for listed zones are added together."
          />
          <FormField
            value={metricName || ''}
            labelWidth={8}
            inputWidth={20}
            onChange={this.onMetricNameChange}
            label="Metric Name"
            tooltip="Graphed metric's name. If empty, a name is generated."
          />
        </div>
      </div>
    );
  }
}
