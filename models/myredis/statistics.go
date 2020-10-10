package myredis

import (
    "deluxe-gin/monitor"
    "github.com/prometheus/client_golang/prometheus"
)

var icount = &monitor.Metric{
    ID:          "myredis_request_count",
    Name:        "redis_request_count",
    Description: "Redis requests detail, partitioned by function and error.",
    Type:        "counter_vec",
    Args:        []string{ "func", "errmsg" },
}

var iduration = &monitor.Metric{
    ID:          "redis_request_duration",
    Name:        "redis_request_duration",
    Description: "The Redis request latencies in milliseconds.",
    Type:        "histogram_vec",
    Args:        []string{ "func" },
}

func GetMetrics() []*monitor.Metric{
    return []*monitor.Metric{ icount, iduration }
}

var count *prometheus.CounterVec
var duration *prometheus.HistogramVec

// 上报调用次数和结果
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportCount( function string, err string ){
    if count == nil{
        if icount.MetricCollector == nil{
            return
        }
        count = icount.MetricCollector.(*prometheus.CounterVec)
    }

    count.WithLabelValues( function, err ).Inc()
}

// 上报调用时延
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportDuration( function string, durMillisecond float64 ){
    if duration == nil{
        if iduration.MetricCollector == nil {
            return
        }
        duration = iduration.MetricCollector.(*prometheus.HistogramVec)
    }

    duration.WithLabelValues( function ).Observe( durMillisecond )
}

