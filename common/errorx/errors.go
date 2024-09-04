package errorx

import (
	"fmt"
	"strings"
)

/**
常用通用固定错误
*/

type CodeError struct {
	code    uint32
	message string
}

//返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.code
}

//返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	if !strings.Contains(e.message, "rpc error") {
		return e.message
	} else {
		// 过滤RPC错误，只显示rpc错误信息
		var msgStrs = strings.Split(e.message, "desc = ")
		return msgStrs[1]
	}
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("Code:%d，Message:%s", e.code, e.message)
}

func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{code: errCode, message: errMsg}
}
func NewErrCode(errCode uint32) *CodeError {
	return &CodeError{code: errCode, message: MapErrMsg(errCode)}
}

func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{code: ServerCommonError, message: errMsg}
}
