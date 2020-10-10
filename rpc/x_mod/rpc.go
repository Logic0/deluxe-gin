package x_mod

import (
    "context"
    "errors"
    "fmt"
    "net/http"

    "deluxe-gin/errno"
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
        ReportCount(appid, cmd, string(rsp.StatusCode()) )
        logx.Errorf("Request to %s failed as status code %d", path, rsp.StatusCode() )
        return nil, errors.New("request x_mod failed" )
    }

    ReportCount(appid, cmd, "ok" )
    logx.Infof("TSPAccess Response BODY: %+v", response )
    return &response, nil
}

// FIXME: 为测试而写的假函数
func FakeExecute( ctx context.Context, factoryID uint, cmd string, body interface{} ) (*Response, error) {
    var response Response
    path := fmt.Sprintf( "/api/v1/x_mod/vehicle/%s?factory_id=%d", cmd, factoryID )
    fmt.Println( "path to tsp access " + path )
    response.Code = errno.SUCCESS
    response.Message = "ok"
    response.Data.SessionID = "123456767899"
    return &response, nil
}

