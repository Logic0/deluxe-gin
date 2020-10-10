package vehicle_api
import (
    "fmt"

    "deluxe-gin/monitor"
    "github.com/prometheus/client_golang/prometheus"
)

var icount = &monitor.Metric{
    ID:          "vehicle_control_request_count",
    Name:        "vehicle_control_request_count",
    Description: "Vehicle Control requests detail, partitioned by function and error.",
    Type:        "counter_vec",
    Args:        []string{ "appid", "func", "errno" },
}

var iduration = &monitor.Metric{
    ID:          "vehicle_control_request_duration",
    Name:        "vehicle_control_request_duration",
    Description: "The Vehicle Control request latencies in milliseconds.",
    Type:        "histogram_vec",
    Args:        []string{ "appid", "func" },
}

func GetMetrics() []*monitor.Metric{
    return []*monitor.Metric{ icount, iduration }
}

var count *prometheus.CounterVec
var duration *prometheus.HistogramVec

// 上报调用次数和结果
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportCount( appid uint, function string, errCode int ){
    if count == nil{
        if icount.MetricCollector == nil{
            return
        }
        count = icount.MetricCollector.(*prometheus.CounterVec)
    }

    count.WithLabelValues( fmt.Sprintf("%d",appid), function, fmt.Sprintf("%d",errCode) ).Inc()
}

// 上报调用时延
// 该函数必须在 monitor.StartStatisticsWithMetrics 调用后才可使用
func ReportDuration( appid uint, function string, durMillisecond float64 ){
    if duration == nil{
        if iduration.MetricCollector == nil {
            return
        }
        duration = iduration.MetricCollector.(*prometheus.HistogramVec)
    }

    duration.WithLabelValues( fmt.Sprintf("%d",appid), function ).Observe( durMillisecond )
}


