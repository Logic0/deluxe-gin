package midware

import (
    "strconv"
    "strings"

    "deluxe-gin/common/proto_common"
    "deluxe-gin/errno"
    log "deluxe-gin/logger"
    "github.com/gin-gonic/gin"
)

func AdminParseAppid() gin.HandlerFunc{
    return func(c *gin.Context) {
        var appid string
        logx := log.SetTraceID( c.Request.Context() )
        appid = c.Query( "appid" )
        appid = strings.Trim(appid, " " )
        if len(appid) < 1 || len(appid) > 10 {
            logx.Errorf( "invalid appid %s", appid )
            proto_common.ErrorResponse( c, errno.NO_PERMISSION, errno.NO_PERMISSION_MSG )
            c.Abort()
            return
        }

        if n, err := strconv.Atoi(appid); err != nil {
            logx.Errorf( "appid %s convert to int failed %s", appid )
            proto_common.ErrorResponse( c, errno.NO_PERMISSION, errno.NO_PERMISSION_MSG )
            c.Abort()
            return
        }else{
            c.Set( "appid", n )
            c.Next()
        }
    }
}
