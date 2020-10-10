package mysql

import (
    "context"
    "errors"
)

// 根据 appid 和 vehicleID 检查当前app 是否具有 abilityID 指定的操作权限
func AppCheckEventValid( ctx context.Context, appid uint, vehicleID uint, eventID uint ) (bool, error){
    app := &Application{
        Appid:appid,
    }

    err := db.Find(app).Related(&app.Events, "Events").Error
    if err != nil {
        return false, errors.New( err.Error() )
    }

    for _,ev := range app.Events {
        if ev.ID == eventID {
            return true, nil
        }
    }

    return false, nil
}

// 查询指定 app+vehicle 的事件集
func AppEventGetByVehicleType( ctx context.Context, appid uint, vehicleID uint, abilityID uint, events *[]Event ) error{
    app := &Application{
        Appid:appid,
    }

    err := db.Find(app).Related(events, "Events").Error
    if err != nil {
        return err
    }

    return nil
}

// 给指定 app + vehicle 添加一项事件接收
func AppEventAdd( ctx context.Context, appid uint, vehicleID uint, eventID uint ) error{
    app := &Application{
        Appid:appid,
    }

    err := db.Find(app).Association("Events").Append(&Event{ID:eventID}).Error
    if err != nil {
        return errors.New( err.Error() )
    }

    return nil
}

// 删除指定 app+vehicle 的一项事件接收
func AppEventDelete( ctx context.Context, appid uint, vehicleID uint, eventID uint ) error{
    app := &Application{
        Appid:appid,
    }

    err := db.Find(app).Association("Events").Delete(&Event{ID:eventID}).Error
    if err != nil {
        return errors.New( err.Error() )
    }

    return nil
}






