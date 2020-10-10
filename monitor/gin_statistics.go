package monitor

import "github.com/gin-gonic/gin"

func StartStatisticsWithMetrics( e *gin.Engine, metrics ...[]*Metric ) error {
    var metricList []*Metric

    for _, ms := range metrics{
        for _,m := range ms{
            metricList = append( metricList, m )
        }
    }

    p := NewPrometheus("deluxe-gin", metricList )
    p.Use( e )
    return nil
}

