package main // import "github.com/regiohelden/innovazammad"

import (
	"encoding/json"
	"log"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stevenroose/gonfig"
)

// defaults
var conf = struct {
	ConfPath string `id:"confpath" short:"c" desc:"path to config file"`
	LogLevel string `default:"debug" short:"l" desc:"log level (debug/info/warn/error)"`
	PBX      struct {
		URL            string    `desc:"URL to innovaphone PBX"`
		EndpointPath   string    `default:"/PBX0/user.soap" desc:"path on innovaphone PBX under which the API is accessible"`
		User           string    `desc:"user for authentication with the PBX"`
		Pass           string    `desc:"password for authentication with the PBX"`
		AppName        string    `default:"zammad_bridge" desc:"application name used when accessing PBX"`
		MonitorUser    string    `desc:"PBX user used for monitoring calls; usually a 'trunk line'"`
		FilterOnGroup  string    `desc:"only process calls for users in this group (if not provided, will submit all calls to zammad)"`
		GroupCacheTime *duration `default:"60s" desc:"time to cache group membership information (used by group filtering); setting to 0 causes a group membership query for each call state change"`
	}
	Zammad struct {
		EndpointURL   string `desc:"URL to Zammad's 'CTI (generic)' endpoint"`
		TrimFirstZero bool   `default:"true" desc:"whether to trim the first zero to normalize submitted phone numbers"`
	}
}{}

func main() {
	err := gonfig.Load(&conf, gonfig.Conf{
		ConfigFileVariable:  "confpath",
		FileDefaultFilename: "innozammad.yaml",
		EnvPrefix:           "INNOVAZAMMAD_",
	})
	if err != nil {
		logrus.Fatalf("could not parse config: %s", err)
	}

	logLevel, err := logrus.ParseLevel(conf.LogLevel)
	if err != nil {
		logrus.Fatalf("could not parse loglevel: %s", err)
	}
	logrus.SetLevel(logLevel)

	confJSON, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	logrus.Debug(string(confJSON))

	// sanity-check options
	// see https://github.com/stevenroose/gonfig/issues/23
	switch "" {
	case conf.PBX.URL:
		logrus.Fatal("must provide PBX URL")
	case conf.PBX.User:
		logrus.Fatal("must provide PBX auth user")
	case conf.PBX.Pass:
		logrus.Fatal("must provide PBX auth password")
	case conf.PBX.MonitorUser:
		logrus.Fatal("must provide PBX monitor user")
	case conf.Zammad.EndpointURL:
		logrus.Fatal("must provide Zammad CTI Endpoint")
	}

	logrus.Infof("starting innovaphone-zammad bridge")

	innovaphoneService := newInnovaphoneService()
	for {
		innovaphoneService.tryConnectionInit()
		if err := innovaphoneService.pollForever(); err != nil {
			logrus.Errorf("error while polling: %s", err)
		}
	}
}

type duration time.Duration

func (d *duration) UnmarshalText(data []byte) error {
	dd, err := time.ParseDuration(string(data))
	*d = duration(dd)
	return err
}
