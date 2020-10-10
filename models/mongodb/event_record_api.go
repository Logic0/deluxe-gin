package mongodb

import (
    "context"

    log "deluxe-gin/logger"
)

// 车辆事件记录
func EventRecordSave( ctx context.Context, r *EventRecord ) error {
    logx := log.SetTraceID( ctx )
    session := mongoSession.Clone()
    defer session.Close()
    collection := session.DB( EVENT_RECORD_DB ).C( EVENT_RECORD_COLLECTION )
    logx.Infof("start to write to mongo: %+v", *r )
    err := collection.Insert( r )
    if err != nil {
        logx.Errorf( "mongodb insert failed as %s", err.Error() )
        return err
    }

    return nil
}

func EventRecordFilter( cond EventRecordFilterCondition, r *[]EventRecord ) error {
    return nil
}
