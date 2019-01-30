package main // import "github.com/regiohelden/innovazammad"

import (
	"expvar"
	"fmt"
	"net/http"
	"os"

	"github.com/regiohelden/innovazammad/config"
	"github.com/regiohelden/innovazammad/innovaphone"
	"github.com/regiohelden/innovazammad/zammad"
	"github.com/sirupsen/logrus"
	"github.com/stevenroose/gonfig"
)

var version = "dev"

func init() {
	err := gonfig.Load(&config.Global, gonfig.Conf{
		ConfigFileVariable:  "confpath",
		FileDefaultFilename: "/etc/innovazammad.yaml",
		EnvPrefix:           "INNOVAZAMMAD_",
	})
	if err != nil {
		logrus.Fatalf("could not parse config: %s", err)
	}
}

func main() {
	if config.Global.Version {
		fmt.Println(version)
		os.Exit(0)
	}
	// used by expvars; exported on localhost because we also publish config (with passwords)
	go http.ListenAndServe("localhost:8080", nil)

	expvar.Publish("config", &config.Global)
	expvar.NewString("version").Set(version)

	logLevel, err := logrus.ParseLevel(config.Global.LogLevel)
	if err != nil {
		logrus.Fatalf("could not parse loglevel: %s", err)
	}
	logrus.SetLevel(logLevel)

	logrus.Debugf("Config: %s", config.Global.String())

	// sanity-check options
	// see https://github.com/stevenroose/gonfig/issues/23
	switch "" {
	case config.Global.PBX.URL:
		logrus.Fatal("must provide PBX URL")
	case config.Global.PBX.User:
		logrus.Fatal("must provide PBX auth user")
	case config.Global.PBX.Pass:
		logrus.Fatal("must provide PBX auth password")
	case config.Global.PBX.MonitorUser:
		logrus.Fatal("must provide PBX monitor user")
	case config.Global.Zammad.EndpointURL:
		logrus.Fatal("must provide Zammad CTI Endpoint")
	}

	logrus.Infof("starting innovaphone-zammad bridge")

	zammadSession := zammad.NewSession()
	for {
		innovaphoneSession := innovaphone.NewSession()
		calls, errs := innovaphoneSession.PollForever()
	handling:
		for {
			select {
			case call := <-calls:
				go zammadSession.Submit(call)
			case err := <-errs:
				logrus.Errorf("error while polling: %s", err)
				break handling
			}
		}
	}
}
