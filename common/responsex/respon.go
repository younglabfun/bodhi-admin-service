package responsex

type ResponseSuccessBean struct {
	//RequestUuid string      `json:"requestUuid"`
	Code    uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	//TimeStamp   int64       `json:"timeStamp"`
}
type NullJson struct{}

func Success(message string, data interface{}) *ResponseSuccessBean {
	return &ResponseSuccessBean{200, message, data}
}

type ResponseErrorBean struct {
	Code    uint32 `json:"code"`
	Message string `json:"message"`
}

func Error(errCode uint32, errMsg string) *ResponseErrorBean {
	return &ResponseErrorBean{errCode, errMsg}
}
