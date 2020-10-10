package sys_admin

import (
    "net/http"

    "deluxe-gin/common/proto_common"
    "deluxe-gin/errno"
    "deluxe-gin/models/mysql"
    "github.com/gin-gonic/gin"
)

// 读全集
func EventGetAll( c *gin.Context ){
    var events []mysql.Event
    err := mysql.EventGetAll( c.Request.Context(), &events )
    if err != nil{
        proto_common.ErrorResponse( c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG )
        return
    }

    var response EventGetAllResponse
    for _, item := range events{
        response.Data.Events = append( response.Data.Events, EventItem{ID: item.ID, Name: item.Name, Desc: item.Desc})
    }

    response.SetError( errno.SUCCESS, errno.SUCCESS_MSG )
    c.JSON( http.StatusOK, &response )
    return
}

// 新增
func EventAdd( c *gin.Context ){
    var request EventAddRequest
    if err := c.ShouldBindJSON( &request ); err != nil{
        proto_common.ErrorResponse(c, errno.FORMAT_ERROR, errno.FORMAT_ERROR_MSG )
        return
    }
    if len( request.Name ) < 1 || len( request.Desc ) < 1{
        proto_common.ErrorResponse( c, errno.PARAM_ERROR, errno.PARAM_ERROR_MSG + "name and desc must not empty")
        return
    }

    if exist, err := mysql.EventExists( c.Request.Context(), request.ID ); err != nil {
        proto_common.ErrorResponse(c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG)
        return
    } else {
        if exist {
            proto_common.ErrorResponse(c, errno.ALREADY_EXIST, errno.ALREADY_EXIST_MSG)
            return
        }
    }

    ev := &mysql.Event{ID:request.ID, Name:request.Name, Desc:request.Desc}
    err := mysql.EventCreate( c.Request.Context(), ev )
    if err != nil{
        proto_common.ErrorResponse( c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG + ":" + err.Error() )
        return
    }

    var response EventAddResponse
    response.SetError( errno.SUCCESS, errno.SUCCESS_MSG )
    c.JSON( http.StatusOK, &response )
    return
}

// 更新
func EventUpdate( c *gin.Context ) {
    var request EventUpdateRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        proto_common.ErrorResponse(c, errno.FORMAT_ERROR, errno.FORMAT_ERROR_MSG)
        return
    }

    if request.ID == 0 || (len(request.Name) < 1 && len(request.Desc) < 1) {
        proto_common.ErrorResponse(c, errno.PARAM_ERROR, errno.PARAM_ERROR_MSG)
        return
    }

    if exist, err := mysql.EventExists( c.Request.Context(), request.ID ); err != nil {
        proto_common.ErrorResponse(c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG)
        return
    } else {
        if !exist {
            proto_common.ErrorResponse(c, errno.NOT_EXIST, errno.NOT_EXIST_MSG)
            return
        }
    }

    if err := mysql.EventUpdate( c.Request.Context(), request.ID, request.Name, request.Desc ); err != nil {
        proto_common.ErrorResponse(c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG+err.Error())
        return
    }

    proto_common.SuccessResponse(c)
    return
}

// 删除
func EventDelete( c *gin.Context ){
    var request EventDeleteRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        proto_common.ErrorResponse(c, errno.FORMAT_ERROR, errno.FORMAT_ERROR_MSG)
        return
    }

    if request.ID == 0 {
        proto_common.ErrorResponse(c, errno.PARAM_ERROR, errno.PARAM_ERROR_MSG)
        return
    }

    if exist, err := mysql.EventExists( c.Request.Context(), request.ID ); err != nil {
        proto_common.ErrorResponse(c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG)
        return
    } else {
        if !exist {
            proto_common.ErrorResponse(c, errno.NOT_EXIST, errno.NOT_EXIST_MSG)
            return
        }
    }

    if err := mysql.EventDelete( c.Request.Context(), request.ID ); err != nil {
        proto_common.ErrorResponse(c, errno.SYSTEM_ERROR, errno.SYSTEM_ERROR_MSG+err.Error())
        return
    }

    proto_common.SuccessResponse(c)
    return
}

// 条件筛选
func EventConditionFilter( c *gin.Context ){

}

