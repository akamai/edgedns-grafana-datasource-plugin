# Build Instructions
Build, package and release the "Edge DNS Datasource" plugin.

## Clone or fork the repository
See [Git Handbook](https://docs.github.com/en/get-started/using-git/about-git) for instructions.  
If you clone, all changes should be made on the 'develop' branch.

## Update the version number
Edit package.json

Advance the version number.  For example:
```
  "version": "2.1.0",
```

## Build
See these references:  
* [Build a plugin](https://grafana.com/developers/plugin-tools)
* [Build a data source plugin](https://grafana.com/developers/plugin-tools/tutorials/build-a-data-source-plugin)

### First time build
This project uses Yarn 4 (managed by corepack). Enable corepack once, then install:
```
corepack enable
yarn install 
```

### Build the back end
Run this command:
```
mage -v
```

My output (after having previously built), looks like this:
```
$ mage -v
Running dependency: github.com/grafana/grafana-plugin-sdk-go/build.Build.LinuxARM-fm
Running dependency: github.com/grafana/grafana-plugin-sdk-go/build.Build.Darwin-fm
Running dependency: github.com/grafana/grafana-plugin-sdk-go/build.Build.Windows-fm
Running dependency: github.com/grafana/grafana-plugin-sdk-go/build.Build.Linux-fm
Running dependency: github.com/grafana/grafana-plugin-sdk-go/build.Build.LinuxARM64-fm
exec: go build -o dist/gpx_akamai-edgedns-datasource-plugin_linux_arm64 -ldflags -w -s -extldflags "-static" ./pkg
exec: go build -o dist/gpx_akamai-edgedns-datasource-plugin_linux_arm -ldflags -w -s -extldflags "-static" ./pkg
exec: go build -o dist/gpx_akamai-edgedns-datasource-plugin_linux_amd64 -ldflags -w -s -extldflags "-static" ./pkg
exec: go build -o dist/gpx_akamai-edgedns-datasource-plugin_windows_amd64.exe -ldflags -w -s -extldflags "-static" ./pkg
exec: go build -o dist/gpx_akamai-edgedns-datasource-plugin_darwin_amd64 -ldflags -w -s -extldflags "-static" ./pkg
```

### Build the front end
Run this command:
```
yarn build
```

My output (after having previously built), looks like this:
```
$ yarn build
(node:93581) [MODULE_TYPELESS_PACKAGE_JSON] Warning: Module type of file:///.../.config/webpack/webpack.config.ts is not specified and it doesn't parse as CommonJS.
Reparsing as ES module because module syntax was detected. This incurs a performance overhead.
To eliminate this warning, add "type": "module" to /.../package.json.
(Use `node --trace-warnings ...` to show where the warning was created)
assets by path *.md 9.95 KiB
  asset README.md 8.89 KiB [emitted] [from: ../README.md] [copied]
  asset CHANGELOG.md 1.06 KiB [emitted] [from: ../CHANGELOG.md] [copied]
asset module.js 13.6 KiB [emitted] [minimized] (name: module) 1 related asset
asset LICENSE 9.94 KiB [emitted] [from: ../LICENSE] [copied]
asset img/akamai-logo.png 1.72 KiB [emitted] [from: img/akamai-logo.png] [copied]
asset plugin.json 1.27 KiB [emitted] [from: plugin.json] [copied]
cached modules 61.7 KiB (javascript) 1.74 KiB (runtime) [cached] 64 modules
webpack 5.108.4 compiled successfully in 81 ms
```

The `MODULE_TYPELESS_PACKAGE_JSON` warning is expected and harmless — Node's ESM loader can't tell from `package.json` alone whether `.config/webpack/webpack.config.ts` is CommonJS or ESM, guesses correctly, and moves on. It does not affect the build result.

## Commit your changes 
See [Git Handbook](https://docs.github.com/en/get-started/using-git/about-git) for instructions.  
Open a Pull Request.

## Package
Copy the 'dist' directory to 'akamai-edgedns-datasource' and then compress.
```
cp -r dist akamai-edgedns-datasource
zip akamai-edgedns-datasource-2.1.0.zip akamai-edgedns-datasource/ -r
```
'2.1.0' is an example. Use your current plugin version number.

## Release

Navigate to https://github.com/akamai/edgedns-grafana-datasource-plugin.

Log in. (You'll need admin rights.)

Follow the directions in [Managing releases in a repository](https://docs.github.com/en/repositories/releasing-projects-on-github/managing-releases-in-a-repository).  
Tags should start with 'v', followed by the build number.  For example, 'v2.1.0'.  

