package mysql

import (
    "context"
    "errors"

    "github.com/jinzhu/gorm"
)

/****************************************** 全集操作 ***********************************/
// 新建
func EventCreate( ctx context.Context, event *Event) error {
    if event == nil{
        return errors.New("in param nil")
    }

    result := db.Create( event )
    if result.Error != nil{
        return result.Error
    }

    return nil
}

// 读全集
func EventGetAll( ctx context.Context, events *[]Event ) error{
    err := db.Find( events ).Error
    if err != nil{
        return err
    }
    return nil
}

// 判定是否存在
func EventExists( ctx context.Context, id uint ) (bool,error){
    var ev Event
    if result := db.First( &ev, Event{ID: id} ); result.Error != nil{
        if gorm.IsRecordNotFoundError( result.Error ){
            return false, nil
        }else{
            return false, result.Error
        }
    }

    return true, nil
}

// 更新单个事件的描述
func EventUpdate( ctx context.Context, id uint, name string, desc string ) error {
    var event Event
    if id == 0 {
        return errors.New("id must greater than 0")
    }

    // // 判断是否存在
    // result := db.First( &event_listener )
    // if result.Error != nil{
    //     if gorm.IsRecordNotFoundError( result.Error ){
    //         return errors.New("not exist")
    //     }else{
    //         return result.Error
    //     }
    // }

    if len(name) != 0{
        event.Name = name
    }
    if len(desc) != 0{
        event.Desc = desc
    }

    result := db.Model(&Event{ID:id}).Update( &event )
    if result.Error != nil{
        if gorm.IsRecordNotFoundError( result.Error ){
            return errors.New("not exist")
        }else{
            return result.Error
        }
    }

    return nil
}

// 删除
func EventDelete( ctx context.Context, id uint ) error {
    result := db.Model(&Event{}).Delete( Event{ID:id} )
    if result.Error != nil{
        return result.Error
    }

    return nil
}


