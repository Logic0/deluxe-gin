package x_mod

import (
    "fmt"

    "deluxe-gin/monitor"
    "github.com/prometheus/client_golang/prometheus"
)

var icount = &monitor.Metric{
    ID:          "x_mod_request_count",
    Name:        "x_mod_request_count",
    Description: "TSP Access service requests detail, partitioned by function and error.",
    Type:        "counter_vec",
    Args:        []string{ "factory_id", "func", "errmsg" },
}

var iduration = &monitor.Metric{
    ID:          "tsp_access_request_duration",
    Name:        "tsp_access_request_duration",
    Description: "The TSP Access service request latencies in milliseconds.",
    Type:        "histogram_vec",
    Args:        []string{ "factory_id", "func" },
}

func GetMetrics() []*monitor.Metric{
    return []*monitor.Metric{ icount, iduration }
}

var count *prometheus.CounterVec
var duration *prometheus.HistogramVec

// 上报调用次数和结果
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportCount( factoryId uint, function string, err string ){
    if count == nil{
        if icount.MetricCollector == nil{
            return
        }
        count = icount.MetricCollector.(*prometheus.CounterVec)
    }

    count.WithLabelValues( fmt.Sprintf("%d",factoryId), function, err ).Inc()
}

// 上报调用时延
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportDuration( factoryId uint, function string, durMillisecond float64 ){
    if duration == nil{
        if iduration.MetricCollector == nil {
            return
        }
        duration = iduration.MetricCollector.(*prometheus.HistogramVec)
    }

    duration.WithLabelValues( fmt.Sprintf("%d",factoryId), function ).Observe( durMillisecond )
}

