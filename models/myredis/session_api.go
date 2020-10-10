package myredis

import (
    "context"
    "encoding/json"
    "errors"

    "deluxe-gin/config"
    log "deluxe-gin/logger"
    "deluxe-gin/session"
    "github.com/gomodule/redigo/redis"
)

// 带重试的 myredis 操作包装，返回结果为 myredis.Do 执行结果
// TODO: 重试逻辑完善
func executeWithRetry( conn redis.Conn, retryTimes int ,commandName string, args ...interface{}) (reply interface{}, err error){
    return conn.Do(commandName, args...)
}

// 存 session 数据到 myredis
func SessionSave( ctx context.Context, session *Session ) error{
    redisConn := redisConnPool.Get()
    logx := log.SetTraceID( ctx )
    defer func(){
        if err := redisConn.Close(); err != nil{
            ReportCount("Close conn", err.Error() )
            logx.Error("myredis connection return failed: " + err.Error() )
        }
    }()

    c, err := json.Marshal( *session )
    if err != nil{
        return errors.New( "json marshal failed")
    }

    // 存 session
    redisKey := session.SessionID
    _, bErr := executeWithRetry( redisConn, 3, "SET", redisKey , string(c) )
    if bErr != nil{
        ReportCount("SET session", bErr.Error() )
        return errors.New( "myredis op1 failed:" + bErr.Error() )
    }
    ReportCount("SET session", "ok" )

    // 设置超时清记录
    _, bErr = executeWithRetry( redisConn,
        3,
        "EXPIRE",
        redisKey,
        config.Config.Session.ExpireTime / 1000 )

    if bErr != nil{
        ReportCount("EXPIRE session", bErr.Error() )
        return errors.New( "myredis op1 failed:" + bErr.Error() )
    }
    ReportCount("EXPIRE session", "ok" )

    return nil
}

// 获取 session 数据
func SessionGet( ctx context.Context, id string ) ( *Session, error ) {
    redisConn := redisConnPool.Get()
    logx := log.SetTraceID( ctx )
    defer func(){
        if err := redisConn.Close(); err != nil{
            ReportCount("Close conn", err.Error() )
            logx.Error("myredis connection return failed: " + err.Error() )
        }
    }()

    val, bErr := redis.Bytes( executeWithRetry( redisConn,3, "GET", id ) )
    if bErr != nil{
        if bErr == redis.ErrNil{
            ReportCount("GET session", "not found" )
            return nil, nil
        }
        ReportCount("GET session", bErr.Error() )
        return nil, bErr
    }

    ReportCount("GET session", "ok" )

    var session Session
    err := json.Unmarshal( val, &session )
    if err != nil{
        return nil, errors.New( "json unmarshal failed")
    }

    return &session, nil
}

// 获取 session 数据
func SessionExist( ctx context.Context, appid uint, id string ) (bool,error) {
    redisConn := redisConnPool.Get()
    logx := log.SetTraceID( ctx )
    defer func(){
        if err := redisConn.Close(); err != nil{
            ReportCount("Close conn", err.Error() )
            logx.Error("myredis connection return failed: " + err.Error() )
        }
    }()

    myID := session.GenerateSessionID( appid, id )
    ok, err := redis.Bool( redisConn.Do("EXISTS", myID ))
    if err != nil {
        ReportCount("EXISTS session", err.Error() )
        return ok,nil
    }
    ReportCount("EXISTS session", "ok" )

    return ok, err
}

// TODO: 删除 myredis 里的 session 存储
func SessionDelete( ctx context.Context, id string ) error {
    return nil
}
