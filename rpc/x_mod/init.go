package x_mod

import (
    "time"

    "deluxe-gin/config"
    log "deluxe-gin/logger"
    "deluxe-gin/monitor"
    "github.com/go-resty/resty/v2"
)

var httpClient *resty.Client

func init(){
    log.Info("[+] 初始化 x_mod http client ...")

    tracedHC, err := monitor.GetTracedHttpClient()
    if err != nil {
        log.Errorf( "monitor.GetTracedHttpClient failed: %v", err )
        panic("tsp access monitor.GetTracedHttpClient failed!")
    }

    httpClient = resty.NewWithClient( tracedHC )
    httpClient.SetTimeout( time.Duration( config.Config.System.TSPAccessTimeout )* time.Millisecond )
    httpClient.SetHostURL( config.Config.System.TSPAccessAddr )
    httpClient.SetHeader("Content-Type", "application/json")
    httpClient.SetLogger( log.New() )
    httpClient.SetDebug( false )
}

