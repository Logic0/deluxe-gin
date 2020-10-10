package errno

type MyError struct{
    Code int
    Message string
}

func (e *MyError) Set( code int, msg string ){
    e.Code = code
    e.Message = msg
}

func (e *MyError) GetMessage() string {
    return e.Message
}

func (e *MyError) GetCode() int{
    return e.Code
}
