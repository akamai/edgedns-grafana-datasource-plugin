# Akamai Edge DNS Datasource

Use the Akamai Edge DNS datasource plugin to observe Edge DNS metrics.

## Prerequisites

* Grafana 7.0 or newer.
* An Akamai API client with authorization to use the [Edge DNS Traffic By Time API](https://developer.akamai.com/api/core_features/reporting/authoritative-dns-traffic-by-time.html). 
* [Authenticate With EdgeGrid](https://developer.akamai.com/getting-started/edgegrid) provides information to generate the required credentials. Choose the API service named "Reporting API", and set the access level to READ-WRITE.

## Installing Grafana

[Install Grafana](https://grafana.com/docs/grafana/latest/installation/) details the process of installing Grafana on several operating systems.

## Installing this plugin on a local Grafana

* On the [edgedns-grafana-datasource-plugin](https://github.com/akamai/edgedns-grafana-datasource-plugin) GitHub repository, 
under "Releases", select "Grafana datasource for Akamai Edge DNS metrics v1.0.0".

* Copy "akamai-edgedns-datasource-1.0.0.zip" to your computer.  Unzip the archive.

### Linux OSs (Debian, Ubuntu, CentOS, Fedora, OpenSuse, Red Hat)

Configuration file: /etc/grafana/grafana.ini  

Plugin directory: /var/lib/grafana/plugins  
Log directory: /var/log/grafana/

Under the plugin directory, create a directory called 'edgedns-grafana-datasource'.

From the unzipped archive, copy:
* LICENSE
* README.md
* img (directory and its contents)
* module.js
* module.js.LICENSE.txt
* module.js.map
* plugin.json  
to /var/lib/grafana/plugins/edgedns-grafana-datasource

From the unzipped archive, copy one of (as appropriate for your hardware):
* gpx_akamai-edgedns-datasource-plugin_linux_amd64
* gpx_akamai-edgedns-datasource-plugin_linux_arm
* gpx_akamai-edgedns-datasource-plugin_linux_arm64  

to /var/lib/grafana/plugins/edgedns-grafana-datasource

### Macintosh

Configuration file: /usr/local/etc/grafana/grafana.ini  

Plugin directory: /usr/local/var/lib/grafana/plugins  
Log directory: /usr/local/var/log/grafana/

Under the plugin directory, create a directory called 'edgedns-grafana-datasource'.

From the unzipped archive, copy:
* LICENSE
* README.md
* img (directory and its contents)
* module.js
* module.js.LICENSE.txt
* module.js.map
* plugin.json  
to /usr/local/var/lib/grafana/plugins/edgedns-grafana-datasource

From the unzipped archive, copy:
* gpx_akamai-edgedns-datasource-plugin_darwin_amd64  
to /usr/local/var/lib/grafana/plugins/edgedns-grafana-datasource

### Windows

Grafana can be installed into any directory (install_dir).
Configuration file: install_dir\conf  

Plugin directory: install_dir\data\plugins  
Log directory: install_dir\data\log

Under the plugin directory, create a directory called 'edgedns-grafana-datasource'.

From the unzipped archive, copy:
* LICENSE
* README.md
* img (directory and its contents)
* module.js
* module.js.LICENSE.txt
* module.js.map
* plugin.json 
to install_dir\data\plugins\edgedns-grafana-datasource 

From the unzipped archive, copy:
* gpx_akamai-edgedns-datasource-plugin_windows_amd64.exe  
to install_dir\data\plugins\edgedns-grafana-datasource

### Grafana configuration

[Configuration](https://grafana.com/docs/grafana/latest/administration/configuration/) 
describes configuration for each operating system.

* Using a text editor, open the configuration file (as described in [Configuration](https://grafana.com/docs/grafana/latest/administration/configuration/).

* Under the [paths] section header, uncomment "plugins".

* To the right of "plugins =", insert the complete path to the plugin directory.

* Under the [plugins] section header, uncomment "allow_loading_unsigned_plugins".
* To the right of "allow_loading_unsigned_plugins =", add "akamai-edgedns-datasource" (without quotes).

### Restart Grafana
[Restart Grafana](https://grafana.com/docs/grafana/latest/installation/restart-grafana/)
describes how to restart Grafana for each operating system.

Under the log directory for your operating system, in "grafana.log", you should see something similar to:
```
t=2021-03-24T10:31:09-0400 lvl=info msg="Registering plugin" logger=plugins id=akamai-edgedns-datasource
```

[Troubleshooting](https://grafana.com/docs/grafana/latest/troubleshooting/) contains troubleshooting tips.

### Log in to Grafana
[Getting started with Grafana](https://grafana.com/docs/grafana/latest/getting-started/getting-started/)
describes how to log in to Grafana.

## "Akamai Edge DNS Datasource" Configuration

Select Configuration (gear icon) -> Datasources -> Akamai Edge DNS Datasource

In the datasource configuration panel, enter your Akamai credentials.

![Data Source](/static/data-source-config.png)

Create a new dashboard and add a panel.

Select a report to graph.

![Report Selection](/static/report-selection.png)

Enter one or more zone names, separated by commas.  If more than one zone is entered then the selected metric (e.g. hits) for all zones are added together and graphed.

![Zones](/static/zones-config.png)

Metric name is optional. If empty then a metric name is automatically generated.

![Metric Name](/static/metric-name-config.png)

