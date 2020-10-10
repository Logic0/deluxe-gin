package main

import (
    "strconv"

    "deluxe-gin/config" // 必须第一个引入，以保证优先初始化
    "deluxe-gin/models/mongodb"
    "deluxe-gin/models/myredis"
    "deluxe-gin/models/mysql"
    "deluxe-gin/monitor"
    "deluxe-gin/rpc/x_mod"
    "github.com/gin-gonic/gin"

    "deluxe-gin/controller/system/sys_admin"
    log "deluxe-gin/logger"
)

func main() {
    gin.SetMode( config.Config.Server.Mode )
    router := gin.New()
    router.Use( gin.Recovery() )
    router.Use( monitor.TracingMidware(router) )

    monitor.StartStatisticsWithMetrics( router,
                                        mysql.GetMetrics(),
                                        myredis.GetMetrics(),
                                        mongodb.GetMetrics(),
                                        x_mod.GetMetrics())

    // 管理后台控制台
    sysAdminRouter := router.Group( "/sys/v1/admin")
    {
        sysAdminRouter.POST( "/add", sys_admin.EventAdd )
        sysAdminRouter.POST( "/get", sys_admin.EventGetAll )
        sysAdminRouter.POST( "/delete", sys_admin.EventDelete )
        sysAdminRouter.POST( "/update", sys_admin.EventUpdate )
        sysAdminRouter.POST( "/filter", sys_admin.EventConditionFilter )
    }

   // 测试接口
    myTest := router.Group("/test")
    {
        myTest.GET("/ab/create", func(c *gin.Context) {
            err := mysql.AppEventAdd( c.Request.Context(), 10000, 100, 101 )
            if err != nil{
                log.Error("AppAbilityAdd failed as %s", err.Error() )
            }
            c.JSON(200, gin.H{
                "hello": "create",
            })
        })

        myTest.GET("/ping",
            func(c *gin.Context) {
                c.JSON(200, gin.H{
                    "message": "pong",
                })
            })

        myTest.GET("/user/:name/*action", func(c *gin.Context) {
            name := c.Param("name")
            action := c.Param("action")
            message := name + " is " + action
            c.String(200, message)
        })

        myTest.GET("/json", func(c *gin.Context) {
            c.JSON(200, gin.H{
                "hello": "<b>hello</b>",
            })
        })

        myTest.GET("/pjson", func(c *gin.Context) {
            c.PureJSON(200, gin.H{
                "hello": "<b>hello</b>",
            })
        })

    }

    log.Infof("=========== server started at :%d in %s mode", config.Config.Server.Port, config.Config.Server.Mode )
    err := router.Run(":" + strconv.FormatUint( uint64(config.Config.Server.Port), 10 ) ) // listen and serve on 0.0.0.0:7070
    if err != nil{
        log.Fatalf( "start gin server failed as %s", err.Error() )
    }
}
