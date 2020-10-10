package mysql

import log "github.com/sirupsen/logrus"

// GormLogger is a custom logger for Gorm, making it use logrus.
type GormLogger struct{}

// Print handles log events from Gorm for the custom logger.
func (*GormLogger) Print(v ...interface{}) {
    switch v[0] {
    case "sql":
        log.WithFields(
            log.Fields{
                "module":  "gorm",
                "type":    "sql",
                "rows":    v[5],
                "src_ref": v[1],
                "values":  v[4],
            },
        ).Debug(v[3])
    case "log":
        log.WithFields(log.Fields{"module": "gorm", "type": "log"}).Print(v[2])
    }
}
