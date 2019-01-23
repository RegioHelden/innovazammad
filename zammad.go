package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/regiohelden/innovazammad/innovaphone"
	"github.com/sirupsen/logrus"
)

func normalizeNumber(n string) string {
	if conf.Zammad.TrimFirstZero {
		return strings.TrimPrefix(n, "0")
	}
	return n
}

func submitToZammad(call *callInSession) error {
	src, dst := call.GetSourceAndDestination()
	dir := call.GetDirection()
	state := call.GetState()

	content := url.Values{
		"callId":    []string{strconv.Itoa(call.Call)},
		"from":      []string{normalizeNumber(src.E164)},
		"to":        []string{normalizeNumber(dst.E164)},
		"direction": []string{dir.String()},
	}
	switch state {
	// we do not get StateAlert on outgoing calls, so we have to assume StateCallProc is enough
	case innovaphone.StateCallProc:
		content.Set("event", "newCall")
	case innovaphone.StateConnect:
		content.Set("event", "answer")
		// TODO: answeringNumber vs user?
	case innovaphone.StateDisconnectSent, innovaphone.StateDisconnectReceived:
		content.Set("event", "hangup")
		// We cannot set "cause" because we cannot differentiate between an answered and non-answered call: we would have
		// to keep state for each call to be able to do that reliably. Zammad, OTOH, seems to handle the call lifetime well
		// enough without it.
	default:
		logrus.Debugf("%s ignoring state change %s; no equivalent on zammad", call.CallInfo, state)
		return nil
	}

	switch dir {
	case innovaphone.DirectionInbound:
		content.Set("user", dst.Cn)
	case innovaphone.DirectionOutbound:
		content.Set("user", src.Cn)
	}

	logrus.Debugf("%s submitting to Zammad: %v", call.CallInfo, content)
	resp, err := http.PostForm(conf.Zammad.EndpointURL, content)
	if err != nil {
		return errors.Wrap(err, "could not submit to Zammad")
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return fmt.Errorf("could not submit to Zammad: %s", resp.Status)
	}
	return nil
}
