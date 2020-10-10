package monitor

import (
    "context"
    "net/http"

    "github.com/SkyAPM/go2sky"
    skyHTTP "github.com/SkyAPM/go2sky/plugins/http"
    "github.com/gin-gonic/gin"
)

func TracingMidware( engine *gin.Engine ) gin.HandlerFunc {
    if engine == nil {
        return func(c *gin.Context) {
            c.Next()
        }
    }

    return Middleware( engine, myTracer )
}

func GetTracedHttpClient( options ... skyHTTP.ClientOption ) (*http.Client, error ){
    httpClient, err := skyHTTP.NewClient( myTracer, options... )
    if err != nil {
        return nil, err
    }

    return httpClient, nil
}

func GetTraceID( ctx context.Context ) string {
    return go2sky.TraceID( ctx )
}
