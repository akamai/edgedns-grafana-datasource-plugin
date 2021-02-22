# Akamai Edge DNS Datasource

Use the Akamai Edge DNS datasource plugin to observe [Edge DNS](https://www.akamai.com/us/en/products/security/edge-dns.jsp) metrics.

## Prerequisites

* Grafana 7.0 or newer.
* An Akamai API client with authorization to use the [Edge DNS Traffic By Time API](https://developer.akamai.com/api/core_features/reporting/authoritative-dns-traffic-by-time.html). 
[Authenticate With EdgeGrid](https://developer.akamai.com/getting-started/edgegrid) provides information to generate the required credentials. 

## Installation

### Grafana Cloud

If you do not have a [Grafana Cloud](https://grafana.com/cloud) account, you can sign up for one [here](https://grafana.com/cloud/grafana).

Click on the `Install Now` button on the [Akamai Edge DNS page on Grafana.com](https://grafana.com/plugins/akamai-edgedns-datasource/installation). This will automatically add the plugin to your Grafana instance. 

### Installing on a local Grafana
Use the grafana-cli tool to install the Edge DNS datasource plugin from the command line:
```bash
grafana-cli plugins install akamai-edgedns-datasource
```

## Configuration

In the datasource configuration panel, enter your Akamai credentials.

![Data Source](/static/DataSourceConfig.png)

Create a new dashboard and add a panel.

Select a report to graph.

![Report Selection](/static/ReportSelection.png)

Enter one or more zone names, separated by commas.  If more than one zone is entered then the selected metric (e.g. hits) for all zones are added together and graphed.

![Zones](/static/ZonesConfig.png)

Metric name is optional. If empty then a metric name is automatically generated.

![Metric Name](/static/MetricNameConfig.png)

## Development

See [Build a data source backend plugin](https://grafana.com/tutorials/build-a-data-source-backend-plugin/) for general build information.

After the source code has been downloaded and your build environment has been setup, you can build at the command line with these commands:

* yarn install
* mage -v
* yarn build

[Configuration](https://grafana.com/docs/grafana/latest/administration/configuration/) has information about platform-specific configuration.

You may want to edit the Grafana configuration file so that Grafana can find and load your newly-built plugin.

For example:
```
[paths]
# Directory where grafana will automatically scan and look for plugins
;plugins = /var/lib/grafana/plugins
``` 
```
[plugins]
# Enter a comma-separated list of plugin identifiers to identify plugins that are allowed to be loaded even if they lack a valid signature.
;allow_loading_unsigned_plugins =
```

Restart Grafana after each new plug-in build.

