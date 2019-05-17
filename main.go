package main // import "github.com/regiohelden/innovazammad"

import (
	"context"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"time"

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

	connectionStats := expvar.NewMap("connection_stats")

	shutdownCtx := GracefulShutdownContext(context.Background(), GracefulShutdownOpts{
		Timeout: time.Duration(*config.Global.GracePeriod),
	})
	// single use short-circuit var; see 'select' below
	shutdown := shutdownCtx.Done()

	// shutdownCtx signals the zammad side to finish processing existing calls, and it in turn signals the innovaphone
	// side when it's done via callHandlingCtx to stop polling.
	zammadSession, callHandlingCtx := zammad.NewSession()

	for {
		innovaphoneSession := innovaphone.NewSession(callHandlingCtx)
		connectionStats.Set("on_since", timeString(time.Now().Format(time.RFC3339)))
		calls, errs := innovaphoneSession.PollForever()

	handling:
		for {
			select {
			case call := <-calls:
				if err := zammadSession.Submit(shutdownCtx, call); err != nil {
					logrus.WithFields(logrus.Fields{
						"call": call,
					}).Warn(err)
				}
			case err := <-errs:
				logrus.Errorf("error while polling: %s", err)
				connectionStats.Add("errors", 1)
				break handling
			case <-shutdown:
				// do a one-time check if we actually have any calls being handled; otherwise we can just quit
				shutdown = nil
				zammadSession.ShutdownIfEmpty()
			case <-callHandlingCtx.Done():
				return
			}
		}
	}
}

type timeString string

func (ts timeString) String() string {
	return fmt.Sprintf("\"%s\"", string(ts))
}
