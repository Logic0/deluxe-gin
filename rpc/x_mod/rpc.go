package x_mod

import (
    "context"
    "errors"
    "fmt"
    "net/http"

    log "deluxe-gin/logger"
    "github.com/go-resty/resty/v2"
)

func myRequest( ctx context.Context ) *resty.Request{
    return httpClient.R().SetContext(ctx)
}

func Execute( ctx context.Context, appid uint, cmd string, body interface{} ) (*Response, error) {
    var response Response
    logx := log.SetTraceID( ctx )
    path := fmt.Sprintf( "/api/v1/x_mod/%s?appid=%d", cmd, appid)
    logx.Debugf("Request BODY: %+v", body )
    rsp, err := httpClient.R().SetContext(ctx).SetBody( body ).SetResult( &response ).Post( path )
    if err != nil{
        ReportCount(appid, cmd, err.Error() )
        logx.Errorf("Request to %s failed as %s", path, err.Error() )
        return nil, errors.New("request x_mod failed" )
    }

    if  rsp.StatusCode() != http.StatusOK{
        ReportCount(appid, cmd, string(rune(rsp.StatusCode())) )
        logx.Errorf("Request to %s failed as status code %d", path, rsp.StatusCode() )
        return nil, errors.New("request x_mod failed" )
    }

    ReportCount(appid, cmd, "ok" )
    logx.Infof("TSPAccess Response BODY: %+v", response )
    return &response, nil
}

