package mysql
// 车型事件的 ID 是固定的
// 为保证预留, 每个事件类型分配 5 个操作控制预留位
const(
    // 电池
    EVENT_1 = 1000
    EVENT_2 = 1001
)

// 有哪些事件，比如 车辆告警、胎压异常、车辆移动、被盗等等
type Event struct{
    ID uint                    `gorm:"PRIMARY_KEY;UNIQUE_INDEX;AUTO_INCREMENT:false"`
    Name string                `gorm:"varchar(60)"`
    Desc string                `gorm:"varchar(255)"`
    Applications []Application `gorm:"many2many:application_event;association_autoupdate:false"`
}

func EventInitData(){
    var ev Event
    db.FirstOrCreate( &ev, Event{ID: EVENT_1, Name:"异常1", Desc:"异常1"} )
    db.FirstOrCreate( &ev, Event{ID: EVENT_2, Name:"异常2", Desc:"异常2"} )

}
