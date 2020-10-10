package mongodb

import "github.com/globalsign/mgo/bson"

type APICallRecord struct{
    ID            bson.ObjectId `bson:"_id,omitempty"`
    Appid         uint          `bson:"appid"`                 //
    OperationTime int64         `bson:"operation_time"`        // 什么时间
    Item          string        `bson:"item"`                  // 操作了那个控制项
    SubItem       int           `bson:"sub_item"`              // 操作了控制项的那个细项
    Action        string        `bson:"action"`                // 是什么操作
}

/***************************** 事件记录 ************************/

const(
    EVENT_NOTIFICATION_SUCCESS      = 1
    EVENT_NOTIFICATION_SYSTEM_ERROR = 2
    EVENT_NOTIFICATION_APP_FAILED   = 3
)

type NotifyResult struct{
    Appid uint                        `bson:"appid"`
    Code int                          `bson:"code"`
    Message string                    `bson:"message"`
}

type EventRecord struct{
    ID bson.ObjectId                  `bson:"_id,omitempty"`
    Appid uint                        `bson:"appid"`
    TriggerTime int64                 `bson:"trigger_time"`
    EventData interface{}             `bson:"event_data"`
    NotifyResult NotifyResult         `bson:"notify_result"`
}

