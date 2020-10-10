package logger

import (
    "fmt"
    "path"
    "runtime"

    "deluxe-gin/config"
    "github.com/sirupsen/logrus"
    "gopkg.in/natefinch/lumberjack.v2"
)

func init(){
    logrus.SetReportCaller( true )
    formatter := &logrus.TextFormatter{
        DisableColors: true,
        CallerPrettyfier: func(f *runtime.Frame) (string, string) {
            filename := path.Base(f.File)
            return fmt.Sprintf("%s()", f.Function), fmt.Sprintf("%s:%d", filename, f.Line)
        },
    }
    logrus.SetFormatter( formatter )

    // 支持日志文件滚动
    logrus.SetOutput(&lumberjack.Logger{
        Filename:   config.Config.Log.Path + config.Config.Log.Filename,
        MaxSize:    config.Config.Log.MaxSize,
        MaxBackups: config.Config.Log.BackupAmount,
        MaxAge:     config.Config.Log.MaxAge,
        Compress:   config.Config.Log.Compress,
    })

    switch config.Config.Log.Level{
    case "debug":
        logrus.SetLevel( logrus.DebugLevel )
    case "info":
        logrus.SetLevel( logrus.InfoLevel )
    case "warn":
        logrus.SetLevel( logrus.WarnLevel )
    case "error":
        logrus.SetLevel( logrus.ErrorLevel )
    case "fatal":
        logrus.SetLevel( logrus.FatalLevel )
    default:
        logrus.SetLevel( logrus.InfoLevel )
    }
}
