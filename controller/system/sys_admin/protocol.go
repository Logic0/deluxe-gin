package sys_admin

import (
    "deluxe-gin/common/proto_common"
)

type AbilityItem struct{
    ID uint                            `json:"id"`
    Name string                        `json:"name"`
    Desc string                        `json:"desc"`
}

// 初始化导入能力全集请求协议
type AbilityImportRequest struct{
    Abilities []AbilityItem            `json:"abilities"`
}

type AbilityImportResponse struct {
    proto_common.CommResponse
}

// 添加一项能力
type AbilityAddRequest struct{
    AbilityItem
}

type AbilityAddResponse struct{
    proto_common.CommResponse
    Data AbilityItem                   `json:"data"`
}

// 更新单项能力集的描述
type AbilityUpdateRequest struct{
    AbilityItem
}

type AbilityUpdateResponse struct{
    proto_common.CommResponse
}

// 删除指定能力
type AbilityDeleteRequest struct{
    ID uint                             `json:"id"`
}

type AbilityDeleteResponse struct{
    proto_common.CommResponse
}

// 获取能力全集的协议
type AbilityGetAllRequest struct{
}

type AbilityGetAllData struct{
    Abilities []AbilityItem            `json:"abilities"`
}

type AbilityGetAllResponse struct{
    proto_common.CommResponse
    Data AbilityGetAllData             `json:"data"`
}

/***************************************** 事件协议 ************************************/
type EventItem struct{
    ID uint                            `json:"id"`
    Name string                        `json:"name"`
    Desc string                        `json:"desc"`
}

// 获取事件全集的协议
type EventGetAllRequest struct{
}

type EventGetAllData struct{
    Events []EventItem                 `json:"events"`
}

type EventGetAllResponse struct{
    proto_common.CommResponse
    Data EventGetAllData               `json:"data"`
}

// 向事件全集添加一个事件
type EventAddRequest struct{
    EventItem
}

type EventAddResponse struct{
    proto_common.CommResponse
}

// 根据条件查询事件
type EventConditionFilterRequest struct{
    Appid uint                          `json:"factory_id"`
    VehicleTypeID uint                  `json:"vehicle_type_id"`
}

type EventGetData struct{
    Events []EventItem                  `json:"events"`
}

type EventConditionFilterResponse struct{
    proto_common.CommResponse
    Data EventGetData                   `json:"data"`
}

// 删除全集里的一个事件
type EventDeleteRequest struct{
    ID uint                             `json:"id"`
}

type EventDeleteResponse struct{
    proto_common.CommResponse
}

// 更新单项事件的描述
type EventUpdateRequest struct{
    EventItem
}

type EventUpdateResponse struct{
    proto_common.CommResponse
}

