# Snap publisher plugin - Warp10

This plugin publishes metrics to Warp10

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Task Manifest Config](#task-manifest-config)
  * [Examples](#examples)
3. [License](#license)
  

### System Requirements
* The Snap daemon is running
* A running version of [Warp10](https://github.com/cityzendata/warp10-platform) reachable by the Snap daemon is required for this plugin to successfully publish data or metrics account at OVH.(https://www.ovh.com/fr/data-platforms/metrics/)

### Installation
#### To build the plugin binary:
This will build the plugin binary in your $GOPATH/bin
```
$ go get github.com/gregleroux/snap-plugin-publisher-warp10
```

Build the plugin by running `make` in the repo:
```
$ make
```
This builds the plugin in `./build`

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation

### Task Manifest Config
A Task Manifest that includes the publishing to warp10 will require configuration data in order for the plugin to establish a connection. Config arguments include:
* "warp_url" (required) - the adress of warp10 host.
* "token" (required) - warp10 token .

### Examples

```

Create a [task manifest](https://github.com/intelsdi-x/snap/blob/master/docs/TASKS.md) (see [exemplary tasks](examples/tasks/)),
for example `psutil-warp10-simple.json` with following content:
```json
{
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "30s"
    },
    "max-failures":10,
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/psutil/load/load1": {},
                "/intel/psutil/load/load5": {},
                "/intel/psutil/load/load15": {},
                "/intel/psutil/cpu/cpu-total/user": {},
                "/intel/psutil/cpu/cpu-total/iowait": {},
                "/intel/psutil/cpu/cpu-total/system": {},
                "/intel/procfs/meminfo/mem_used": {}
            },
          },
          "publish": [
              {
              "plugin_name": "warp10",
                  "config": {
                      "warp_url": "https://127.0.0.1/api/v0/update",
                      "token": "xxxxxxxxx"
                      }
              }
          ]
        }
    }
}

```

## License
[License](LICENSE)

