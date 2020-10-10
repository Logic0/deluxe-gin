package mysql

import (
    "fmt"

    "deluxe-gin/config"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/sirupsen/logrus"
)

var db *gorm.DB

func init() {
    var err error
    dsn := fmt.Sprintf( "%s&timeout=%dms&readTimeout=%dms&writeTimeout=%dms",
        config.Config.MySQL.Uri,
        config.Config.MySQL.ConnectTimeout,
        config.Config.MySQL.ReadTimeout,
        config.Config.MySQL.WriteTimeout )
    logrus.Info("[+] 链接数据库...")
    logrus.Infof("   [-] 数据库: %s", dsn )
    db, err = gorm.Open("mysql", dsn )
    if err != nil {
        logrus.Fatalf("failed to open database: %s", err.Error() )
        panic("failed to open database" )
    }

    db.LogMode(true)
    db.SetLogger( &GormLogger{} )
    db.DB().SetMaxIdleConns(config.Config.MySQL.MaxIdleConn)
    db.DB().SetMaxOpenConns(config.Config.MySQL.MaxIdleConn + config.Config.MySQL.MaxActiveConn)

    // 自动生成表
    logrus.Info("[+] 初始化数据库表结构...")
    db.AutoMigrate( &Event{}, &Application{} )
    logrus.Info("[+] 初始化数据库成功")
}
