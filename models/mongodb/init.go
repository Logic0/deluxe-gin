package mongodb

import (
    "fmt"
    "time"

    "deluxe-gin/config"
    log "deluxe-gin/logger"
    "github.com/globalsign/mgo"
)

const(
    API_RECORD_DB = "api_record"
    API_RECORD_COLLECTION = "call"

    EVENT_RECORD_DB = "event_record"
    EVENT_RECORD_COLLECTION = "alarm"
)

var mongoSession *mgo.Session

func init(){
    dialInfo := &mgo.DialInfo{
        Addrs:config.Config.MongoDB.Servers,
        Timeout: time.Duration( config.Config.MongoDB.ConnectTimeout ) * time.Millisecond,
        ReadTimeout: time.Duration( config.Config.MongoDB.ReadTimeout ) * time.Millisecond,
        WriteTimeout: time.Duration(config.Config.MongoDB.WriteTimeout) * time.Millisecond,
        Username: config.Config.MongoDB.Username,
        Password: config.Config.MongoDB.Password,
        Database: config.Config.MongoDB.AuthDatabase,
        PoolLimit: 100,                               // 单机链接池大小
        PoolTimeout: time.Duration(1) * time.Second,  // 取链接最长等待时间
    }

    log.Info("[+] MongoDB连接中...")
    var err error
    mongoSession, err = mgo.DialWithInfo( dialInfo )
    if err != nil{
        log.Info( "   [-] 链接失败" + err.Error() )
        log.Fatalf( "mongodb connect failed as %s", err.Error() )
        return
    }
    log.Info("[+] MongoDB 初始化成功" )

    log.Info("[+] MongoDB 初始化 api_record/call 索引..." )
    apiIndex := mgo.Index{
        Key: []string{"appid","vehicle_id","session_id"},
    }

    apiCollection := mongoSession.DB(API_RECORD_DB).C(API_RECORD_COLLECTION)
    err = apiCollection.EnsureIndex( apiIndex )
    if err != nil {
        fmt.Println( "  [-] vehicle_api record 索引初始化失败:" + err.Error() )
        log.Errorf("create vehicle_api record index failed as ", err.Error() )
    }
    log.Info("[+] MongoDB 初始化 api_record/call 索引成功" )
}
