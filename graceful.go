package main

import (
	"context"
	"math"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

// DefaultSignals is the default value for GracefulShutdownOpts.Signals
var DefaultSignals = []os.Signal{syscall.SIGTERM, os.Interrupt}

// DefaultGracePeriod is the default value for GracefulShutdownOpts.Timeout
var DefaultGracePeriod = time.Duration(math.MaxInt64)

// GracefulShutdownOpts are options for GracefulShutdownContext
type GracefulShutdownOpts struct {
	// Signals is a list of signals that will trigger a graceful shutdown.
	Signals []os.Signal
	// Timeout is the amount of time to wait between receiving a signal in Signals and exiting.
	// Leaving this at 0 means waiting indefinitely.
	Timeout time.Duration
}

// GracefulShutdownContext returns a context that will be cancelled upon receiving any of the os.Signals in
// GracefulShutdownOpts.Signals.
// After receiving the first signal, we wait for either a second os.Signal or a timer of duartion
// GracefulShutdownOpts.Timeout, whichever comes first, and terminate immediately.
// If GracefulShutdownOpts.Timeout is 0, we wait indefinitely.
func GracefulShutdownContext(ctx context.Context, opts GracefulShutdownOpts) context.Context {
	listenSigs := opts.Signals
	if listenSigs == nil {
		listenSigs = DefaultSignals
	}

	timeout := opts.Timeout
	if timeout == time.Duration(0) {
		timeout = DefaultGracePeriod
	}
	incomingSigs := make(chan os.Signal, 1)
	signal.Notify(incomingSigs, listenSigs...)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		select {
		case got := <-incomingSigs:
			logrus.Infof("received %s, gracefully terminating...", got)
			cancel()
		case <-ctx.Done():
			cancel()
			return
		}

		select {
		case <-incomingSigs:
		case <-time.After(timeout):
			logrus.Warn("grace period over; exiting.")
		}
		os.Exit(1)
	}()

	return ctx
}
