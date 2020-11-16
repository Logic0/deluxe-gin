package config

// 日志配置
type iLog struct{
    Level string              `yaml:"level"`               // 日志输出等级
    Path string               `yaml:"path"`                // 日志文件路径
    Filename string           `yaml:"filename"`            // 日志文件的名字
    MaxAge int                `yaml:"max_age"`             // 最大有效时间, 单位 天
    MaxSize int               `yaml:"max_size"`            // 最大文件容量, 单位 m
    BackupAmount int          `yaml:"backup_amount"`       // 最多文件个数, 单位 个
    Compress bool             `yaml:"compress"`            // 是否对旧文件压缩
    Localtime bool            `yaml:"localtime"`          // 是否使用本地时间, 否则使用 UTC
}

// 监控配置
type iMonitor struct{
    TracingServerAddr   string     `yaml:"tracing_server_addr"`     // 调用链跟踪系统上报地址
    PrometheusExportURL string     `yaml:"prometheus_export_url"`   // prometheus 对外暴露的地址
}

// MySQL
type iMySQL struct{
    Uri string                `yaml:"uri"`
    MaxIdleConn int           `yaml:"max_idle_conn"`
    MaxActiveConn int         `yaml:"max_active_conn"`
    ConnectTimeout int        `yaml:"connect_timeout"`
    ReadTimeout int           `yaml:"read_timeout"`
    WriteTimeout int          `yaml:"write_timeout"`
}

// Redis
type iRedis struct{
    Uri string                `yaml:"uri"`
    MaxIdleConn int           `yaml:"max_idle_conn"`
    MaxActiveConn int         `yaml:"max_active_conn"`
    ConnectTimeout int        `yaml:"connect_timeout"`
    ReadTimeout int           `yaml:"read_timeout"`
    WriteTimeout int          `yaml:"write_timeout"`
}

// MongoDB
type iMongoDB struct{
    Servers []string             `yaml:"servers,flow"`
    AuthDatabase string          `yaml:"auth_database"`
    ConnectTimeout int           `yaml:"connect_timeout"`
    ReadTimeout int              `yaml:"read_timeout"`
    WriteTimeout int             `yaml:"write_timeout"`
    Username string              `yaml:"username"`
    Password string              `yaml:"password"`
}

// 周边系统相关的配置
type iSystem struct{
    XModAddr    string `yaml:"x_mod_addr"`
    XModTimeout int    `yaml:"x_mod_timeout"`
}

// 与 TSP access 交互的 session 配置
type iSession struct{
    ExpireTime int                 `yaml:"expire_time"`
}

// 当前 server 运行配置
type iServer struct{
    Port uint                       `yaml:"port"`
    Mode string                     `yaml:"mode"`
}

type config struct{
    Environment string              `yaml:"-"`             // 区分不同的运行环境, "dev"/"pro"
    MySQL iMySQL                    `yaml:"mysql"`
    Redis iRedis                    `yaml:"redis"`
    MongoDB iMongoDB                `yaml:"mongodb"`
    System iSystem                  `yaml:"system"`
    Session iSession                `yaml:"session"`
    Server iServer                  `yaml:"server"`
    Log iLog                        `yaml:"log"`
    Monitor iMonitor                `yaml:"monitor"`
}
