package mongodb

type TimeRange struct{
    Start int64
    End int64
}

// API 调用记录筛选条件
type APIRecordFilterCondition struct{
    Appid uint                               // 哪个appid 发起的请求
    TimeRange TimeRange                      // 哪个时间段
}

// 事件筛选条件
type EventRecordFilterCondition struct{
    Appid uint                                //
    TimeRange TimeRange                       // 哪个时间段
}
