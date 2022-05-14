# gin-contrib-log

```gin-contrib-log demo
 
package main

import (
	"github.com/donetkit/gin-contrib-log/glog"
	"github.com/sirupsen/logrus"
)

var log glog.ILogger

func main() {
	l := glog.New(glog.WithDisplayFields(true), glog.WithHostName(""))
	l.SetLevel(glog.DebugLevel)
	// enable/disable file/function name
	l.SetReportCaller(false)

	l.Infof("this is %v demo", "TestLogs")

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
	l.Infof("this is %v demo", "TestLogs")
	l.Info("demo end.")
}


```