package main

import (
	"github.com/donetkit/contrib-log/glog"
	"github.com/sirupsen/logrus"
)

var log glog.ILogger

type Person struct {
	name string
	age  int
	sex  string
}

func main() {

	new(logrus.Entry).Info()
	
	l := glog.New(glog.WithDisplayFields(false))
	l.SetLevel(glog.DebugLevel)
	// enable/disable file/function name
	l.SetReportCaller(false)

	lWebServer := l.WithField("component", "web-server")
	lWebServer.Info("starting...")

	lWebServerReq := lWebServer.WithFields(logrus.Fields{
		"req":   "GET /api/stats",
		"reqId": "#1",
	})

	lWebServerReq.Info("params: startYear=2048")
	lWebServerReq.Error("response: 400 Bad Request")

	lDbConnector := l.WithField("category", "db-connector")
	lDbConnector.Info("connecting to db on 10.10.10.13...")
	lDbConnector.Warn("connection took 10s")

}
