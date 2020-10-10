package proto_common

import (
	"net/http"

	"deluxe-gin/errno"
	log "deluxe-gin/logger"
	"github.com/gin-gonic/gin"
)

type CommResponse struct{
	Code int                 `json:"code"`
	Message string           `json:"message"`
}

func (r *CommResponse) SetError( code int, msg string ){
	r.Code = code
	r.Message = msg
}

func ErrorResponse( c *gin.Context, errno int, errmsg string ) {
	var rsp CommResponse
	rsp.SetError( errno, errmsg )
	c.JSON( http.StatusOK, rsp )
	logx := log.SetTraceID( c.Request.Context() )
	logx.Debugf("Response BODY:%+v", rsp )
}

func SuccessResponse( c *gin.Context ){
	var rsp CommResponse
	rsp.SetError( errno.SUCCESS, errno.SUCCESS_MSG )
	c.JSON( http.StatusOK, rsp )
	logx := log.SetTraceID( c.Request.Context() )
	logx.Debugf("Response BODY:%+v", rsp )
}
