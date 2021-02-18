# edgedns-grafana-datasource-plugin

Use Akamai's Edge DNS datasource plugin to observe [Edge DNS](https://www.akamai.com/us/en/products/security/edge-dns.jsp) metrics.

## Prerequisites

* This plugin requires Grafana 7.0 or newer.
* An Akamai API client with authorization to use the [Edge DNS Traffic By Time API](https://developer.akamai.com/api/core_features/reporting/authoritative-dns-traffic-by-time.html). 
[Authenticate With EdgeGrid](https://developer.akamai.com/getting-started/edgegrid) provides an overview and information to generate credentials required to use the API. 

## Installation

### Grafana Cloud

If you do not have a [Grafana Cloud](https://grafana.com/cloud) account, you can sign up for one [here](https://grafana.com/cloud/grafana).

Click on the `Install Now` button on the [Akamai Edge DNS page on Grafana.com](https://grafana.com/plugins/edgedns-grafana-datasource/installation). This will automatically add the plugin to your Grafana instance. 

### Installing on a local Grafana
Use the grafana-cli tool to install the Edge DNS datasource plugin from the command line:
```bash
grafana-cli plugins install edgedns-grafana-datasource-plugin
```

## Configuration

In the Edge DNS datasource configuration panel, enter the Akamai API credentials.

