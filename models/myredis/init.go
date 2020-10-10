package myredis

import (
    "time"

    "github.com/sirupsen/logrus"

    "deluxe-gin/config"
    log "deluxe-gin/logger"
    "github.com/gomodule/redigo/redis"
)

var redisConnPool *redis.Pool

func testOnBorrow( c redis.Conn, t time.Time ) error {
    _, err := c.Do("ping")
    if err != nil {
    }

    return err
}

func dialRedis() (redis.Conn, error){
    logrus.Info("[+] Redis连接池初始化...")
    c, err := redis.DialURL(config.Config.Redis.Uri,
        redis.DialConnectTimeout( time.Duration( config.Config.Redis.ConnectTimeout ) * time.Millisecond),
        redis.DialWriteTimeout( time.Duration(config.Config.Redis.WriteTimeout) * time.Millisecond ),
        redis.DialReadTimeout( time.Duration(config.Config.Redis.ReadTimeout) * time.Millisecond) )

    if err != nil {
        logrus.Warn("[-] Redis连接池初始化失败...")
        log.Errorf("Building myredis connection pool failed: %s", err.Error() )
        panic(err.Error())
    }
    logrus.Info("[+] Redis连接池初始化成功")

    return c, err
}

func init(){
    redisConnPool = &redis.Pool{
        MaxIdle: config.Config.Redis.MaxIdleConn,
        MaxActive: config.Config.Redis.MaxActiveConn,
        IdleTimeout: time.Duration(300) * time.Second,
        Dial: dialRedis,
        TestOnBorrow:testOnBorrow,
    }
}
