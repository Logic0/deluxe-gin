package x_mod

type Request struct{

}

type Data struct{
    SessionID string                              `json:"session_id"`
}

type Response struct{
    Code int               `json:"code"`
    Message string         `json:"message"`
    Data Data              `json:"data"`
}


