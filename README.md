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

| Flag | Description | Default | Required? |
|---|---|---|---|
|`--confpath` | Path to configuration file.  | `/etc/innovazammad.yaml` | |
|`-l`, `--loglevel` | How much logging will be output to stdout (values as supported by [logrus](https://github.com/sirupsen/logrus)). | `warn` | |
|`--graceperiod` | How long to wait for ongoing calls to finish before shutting down or restarting. Restarting may leave orphaned calls on the Zammad side. | `60s` ||
|`--pbx.url` | Under which URL should `innovazammad` attempt to connect to the PBX. | _none_ | × |
|`--pbx.endpointpath` | Path under `pbx.url` where the API is accessible. | `/PBX0/user.soap` ||
|`--pbx.user` | User for authentication with the PBX. | _none_ | × |
|`--pbx.pass` | Password for authentication with the PBX. | _none_ | × |
|`--pbx.appname` | Application name used when connecting to the PBX (informational only). | `innovazammad` ||
|`--pbx.monitoruser` | PBX user object for which events will be monitored. Possibly a 'trunk line'. | _none_ | × |
|`--pbx.filterongroup` | Only events for users in this group will be submitted to Zammad. If not provided, all calls will be submitted. | _none_ ||
|`--pbx.groupcachetime` | Time to cache group membership information used by `pbx.filterongroup`. Setting this to `0` causes group membership to be checked for every incoming event.| `300s` ||
|`--zammad.endpointurl` | URL to Zammad's 'CTI (generic)' endpoint. Can be found in Zammad's integration settings page. | _none_ |×|
|`--zammad.trunkprefix` | Optional prefix to trim from the phone numbers received from the PBX. Normally necessary if calls come through a trunk line. | `0` ||
|`--zammad.countrycode` | Country code to prepend to area-local phone numbers received from the PBX. | `49` ||
|`--zammad.numberprefix` | Optional number to prepend to unqualified phone numbers received from the PBX. This should transform an internal extension into an E123 number. | ||

The environment variable names can be created by prefixing the flag name with `INNOVAZAMMAD_` and replacing any dot (`.`) with underscore (`_`). E.g.: `--pbx.url` becomes `INNOVAZAMMAD_PBX_URL`.

## Compatibility

Currently `innovazammad` uses innovaphone's v11 and newer SOAP API (see `innovaphone/pbx.wsdl`). Older PBX versions might not work.

On the Zammad side, it has been tested against the `2.9`, `3.0` and `3.1` branch.

## Monitoring

Internal state can be monitored by querying the metrics endpoint:
```
curl -s localhost:8080/debug/vars
```
