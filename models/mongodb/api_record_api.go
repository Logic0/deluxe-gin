package mongodb

import (
    "context"

    log "deluxe-gin/logger"
    "github.com/globalsign/mgo"
    "github.com/globalsign/mgo/bson"
)

// 保存 API 调用记录
func APIRecordSave( ctx context.Context, r *APICallRecord) error{
    session := mongoSession.Clone()
    defer session.Close()
    logx := log.SetTraceID( ctx )
    collection := session.DB( API_RECORD_DB ).C( API_RECORD_COLLECTION )
    logx.Infof("start to write to mongo: %+v", *r )
    err := collection.Insert( r )
    if err != nil {
        logx.Errorf( "mongodb insert failed as %s", err.Error() )
        return err
    }

    return nil
}

// TODO: 筛选 API 调用记录
func APIRecordFilter( cond APIRecordFilterCondition, r *[]APICallRecord ) error{
    return nil
}

// 查询 api 执行的结果, sessionID 是端上生成, 有可能会重复, 所以这里的查询结果是一个列表
func APIRecordQuery( ctx context.Context, appid uint, appSessionID string ) (*APICallRecord, error) {
    session := mongoSession.Clone()
    defer session.Close()
    collection := session.DB( API_RECORD_DB ).C( API_RECORD_COLLECTION )
    findM := bson.M{
        "appid": appid,
        "app_session_id": appSessionID,
    }

    var record APICallRecord
    err := collection.Find( findM ).Sort("-operation_time").One( &record )
    if err != nil{
        if err == mgo.ErrNotFound{
            return nil,NOT_FOUND
        }else{
            log.Errorf( "mongo execute failed with %s", err )
            return nil, err
        }
    }

    return &record, nil
}

// 更新 record 执行结果
func APIRecordUpdate( ctx context.Context, r *APICallRecord) (info *mgo.ChangeInfo, err error ){
    session := mongoSession.Clone()
    defer session.Close()
    collection := session.DB( API_RECORD_DB ).C( API_RECORD_COLLECTION )
    findM := bson.M{
        "appid": r.Appid,
    }

    updateM := bson.M{
        "$set":bson.M{
            "operation_time": r.OperationTime,
            "item": r.Item,
            "sub_item": r.SubItem,
            "action": r.Action,
        },
    }

    return collection.Upsert( findM, updateM )
}


