package monitor

import (
	"deluxe-gin/config"
	"github.com/SkyAPM/go2sky"
	"github.com/SkyAPM/go2sky/reporter"
)

var myReporter go2sky.Reporter
var myTracer *go2sky.Tracer

func init() {
	var err error
	switch config.Config.Environment {
	case "dev":
		myReporter, err = reporter.NewLogReporter()
	case "pro":
		myReporter, err = reporter.NewGRPCReporter(config.Config.Monitor.TracingServerAddr)
	default:
		panic("error environment")
	}
	if err != nil {
		panic("Create skywalking reporter failed!")
	}

	myTracer, err = go2sky.NewTracer("deluxe-gin", go2sky.WithReporter(myReporter))
	if err != nil {
		panic("Create skywalking tracer failed!")
	}
}
