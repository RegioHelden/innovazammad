package config

import (
	"encoding/json"
	"fmt"
	"time"
)

// Global is meant to ease access to global configuration, instead of passing it around.
var Global = Config{}

// Config holds the configurable options with their respective defaults.
type Config struct {
	Version     bool      `desc:"display current version"`
	ConfPath    string    `id:"confpath" desc:"path to config file"`
	LogLevel    string    `default:"warn" short:"l" desc:"log level (debug/info/warn/error)"`
	GracePeriod *Duration `desc:"time to wait for ongoing calls during shutdown" default:"60s"`
	PBX         struct {
		URL            string    `desc:"URL to innovaphone PBX"`
		EndpointPath   string    `default:"/PBX0/user.soap" desc:"path on innovaphone PBX under which the API is accessible"`
		User           string    `desc:"user for authentication with the PBX"`
		Pass           string    `desc:"password for authentication with the PBX"`
		AppName        string    `default:"innovazammad" desc:"application name used when accessing PBX"`
		MonitorUser    string    `desc:"PBX user used for monitoring calls; usually a 'trunk line'"`
		FilterOnGroup  string    `desc:"only process calls for users in this group (if not provided, will submit all calls to Zammad)"`
		GroupCacheTime *Duration `default:"300s" desc:"time to cache group membership information (used by group filtering); setting to 0 causes a group membership query for each call state change"`
	}
	Zammad struct {
		EndpointURL   string `desc:"URL to Zammad's 'CTI (generic)' endpoint"`
		TrimFirstZero bool   `default:"true" desc:"whether to trim the first zero to normalize submitted phone numbers"`
		CountryCode   int    `default:"49" desc:"Country code to prepend to unqualified phone numbers received from the PBX"`
		NumberPrefix  string `default:"" desc:"Optional number to prepend to unqualified phone numbers received from the PBX. This should transform an internal extension into an E123 number."`
	}
}

// Used by expvar
func (c *Config) String() string {
	out, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return "ERR"
	}
	return string(out)
}

// Duration wraps time.Duration, adding UnmarshalText for config-parsing
type Duration time.Duration

// UnmarshalText parses the provided byte array into the duration receiver as text
func (d *Duration) UnmarshalText(data []byte) error {
	dd, err := time.ParseDuration(string(data))
	*d = Duration(dd)
	return err
}

// MarshalJSON returns a json-string representation of the duration
func (d *Duration) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Duration(*d).String())), nil
}
