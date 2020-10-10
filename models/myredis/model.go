package myredis

// 存储到 myredis 的 session 结构
type Session struct{
    SessionID     string `json:"session_id"`
    Appid         uint   `json:"appid"`
    Item          string `json:"item"`
    SubItem       string `json:"sub_item"`
}

