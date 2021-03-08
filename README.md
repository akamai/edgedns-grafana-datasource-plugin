# Akamai Edge DNS Datasource

Use the Akamai Edge DNS datasource plugin to observe Edge DNS metrics.

## Prerequisites

* Grafana 7.0 or newer.
* An Akamai API client with authorization to use the [Edge DNS Traffic By Time API](https://developer.akamai.com/api/core_features/reporting/authoritative-dns-traffic-by-time.html). 
* [Authenticate With EdgeGrid](https://developer.akamai.com/getting-started/edgegrid) provides information to generate the required credentials. Choose the API service named "Reporting API", and set the access level to READ-WRITE.

## Installation

### Grafana Cloud

If you do not have a [Grafana Cloud](https://grafana.com/cloud) account, you can sign up for one [here](https://grafana.com/cloud/grafana).

Click on the `Install Now` button on the [Akamai Edge DNS page on Grafana.com](). This will automatically add the plugin to your Grafana instance. 

### Installing on a local Grafana
Use the grafana-cli tool to install the Edge DNS datasource plugin from the command line:
```bash
grafana-cli plugins install akamai-edgedns-datasource
```

## Configuration

In the datasource configuration panel, enter your Akamai credentials.

![Data Source](https://github.com/akamai/edgedns-grafana-datasource-plugin/blob/develop/static/data-source-config.png)

Create a new dashboard and add a panel.

Select a report to graph.

![Report Selection](https://github.com/akamai/edgedns-grafana-datasource-plugin/blob/develop/static/report-selection.png)

Enter one or more zone names, separated by commas.  If more than one zone is entered then the selected metric (e.g. hits) for all zones are added together and graphed.

![Zones](https://github.com/akamai/edgedns-grafana-datasource-plugin/blob/develop/static/zones-config.png)

Metric name is optional. If empty then a metric name is automatically generated.

![Metric Name](https://github.com/akamai/edgedns-grafana-datasource-plugin/blob/develop/static/metric-name-config.png)

