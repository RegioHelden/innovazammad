[![Build Status](https://travis-ci.org/RegioHelden/innovazammad.svg?branch=master)](https://travis-ci.org/RegioHelden/innovazammad)
[![Go Report Card](https://goreportcard.com/badge/github.com/regiohelden/innovazammad)](https://goreportcard.com/report/github.com/regiohelden/innovazammad)

# innovazammad

Forwards call information from an [innovaphone](https://www.innovaphone.com/) PBX to [Zammad](https://zammad.com/)'s generic CTI API.


## Overview

`innovazammad` polls the innovaphone PBX' SOAP API for call events and forwards those to Zammad, using its [generic CTI API](https://docs.zammad.org/en/latest/cti-api-push.html). It keeps track of call status and attempts to map PBX call events to their equivalents on the Zammad side, including keeping track of forwarded calls.

## Installation

Released binaries can be downloaded from the [releases page](https://github.com/RegioHelden/innovazammad/releases).

### From source

To install from source, simply run:
```
go get github.com/regiohelden/innovazammad
```
For go < 1.12 this [will not respect the exact versions used when doing a release](https://github.com/golang/go/issues/24250). To ensure the same results as a released binary for older versions of go, do this instead:
```
git clone https://github.com/regiohelden/innovazammad.git
cd innovazammad
go build
```

## Configuration

Configuration can be supplied as flags, environment variables or in a configuration file (in JSON, YAML or TOML). 

The following options are available:

| Flag | Environment | Description | Default | Required? |
|---|---|---|---|---|
|`--confpath`| `INNOVAZAMMAD_CONFPATH` | Path to configuration file.  | `innovazammad.yaml` | |
|`-l`, `--loglevel`| `INNOVAZAMMAD_LOGLEVEL` | How much logging will be output to stdout (values as supported by [logrus](https://github.com/sirupsen/logrus)). | `warn` | |
|`--pbx.url`| `INNOVAZAMMAD_PBX_URL` | Under which URL should `innovazammad` attempt to connect to the PBX. | _none_ | × |
|`--pbx.endpointpath`| `INNOVAZAMMAD_PBX_ENDPOINTPATH` | Path under `pbx.url` where the API is accessible. | `/PBX0/user.soap` ||
|`--pbx.user`| `INNOVAZAMMAD_PBX_USER` | User for authentication with the PBX. | _none_ | × |
|`--pbx.pass`| `INNOVAZAMMAD_PBX_PASS` | Password for authentication with the PBX. | _none_ | × |
|`--pbx.appname`| `INNOVAZAMMAD_PBX_APPNAME` | Application name used when connecting to the PBX (informational only). | `innovazammad` ||
|`--pbx.monitoruser`| `INNOVAZAMMAD_PBX_MONITORUSER` | PBX user object for which events will be monitored. Possibly a 'trunk line'. | _none_ | × |
|`--pbx.filterongroup`|`INNOVAZAMMAD_PBX_FILTERONGROUP` | Only events for users in this group will be submitted to Zammad. If not provided, all calls will be submitted. | _none_ ||
|`--pbx.groupcachetime`|`INNOVAZAMMAD_PBX_GROUPCACHETIME`| Time to cache group membership information used by `pbx.filterongroup`. Setting this to `0` causes group membership to be checked for every incoming event.| `300s` ||
|`--zammad.endpointurl`|`INNOVAZAMMAD_ZAMMAD_ENDPOINTURL`| URL to Zammad's 'CTI (generic)' endpoint. Can be found in Zammad's integration settings page. | _none_ |×|
|`--zammad.trimfirstzero`| `INNOVAZAMMAD_ZAMMAD_TRIMFIRSTZERO` | Whether to remove a first zero from incoming numbers from the PBX. This will probably be necessary for cases where calls come through a trunk line. | `true` ||

## Compatibility

Currently `innovazammad` uses innovaphone's [v10 SOAP API](http://wiki.innovaphone.com/index.php?title=Reference10:SOAP_API) (see `innovaphone/pbx10_00.{go,wsdl}`). Older PBX might not work.

On the Zammad side, it has been tested against the `2.9.x` branch.