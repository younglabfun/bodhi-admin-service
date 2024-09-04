package errorx

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"                             // 200 grpc code 0 OK
	message[Canceled] = "Canceled"                      // 408 grpc code 1 Canceled 客户端取消访问
	message[Unknown] = "Unknown"                        // 500 grpc code 2 Unknown 未知错误
	message[InvalidArgument] = "Invalid argument"       // 400 grpc code 3 Invalid argument 错误参数
	message[DeadlineExceeded] = "Deadline exceeded"     // 504 grpc code 4 Deadline exceeded 超时访问
	message[NotFound] = "Not found"                     // 404 grpc code 5 Not found 没有找到 非grpc错误
	message[AlreadyExists] = "Already exists"           // 409 grpc code 6 Already exists 非grpc错误
	message[PermissionDenied] = "Permission denied"     // 403 grpc code 7 Permission denied 无权访问，除鉴权中间件以外， 非grpc错误
	message[ResourceExhausted] = "Resource exhausted"   // 429 grpc code 8 Resource exhausted 资源耗尽或空间不足 grpc错误
	message[FailedPrecondition] = "Failed precondition" // 400 grpc code 9 Failed precondition 非grpc错误
	message[Aborted] = "Aborted"                        // 409 grpc code 10 Aborted 终止服务 非grpc错误
	message[OutOfRange] = "Out Of Range"                // 400 grpc code 11 Out Of Range 超出有效范围 非grpc错误
	message[Unimplemented] = "Unimplemented"            // 501 grpc code 12 Unimplemented 未实现方法 grpc错误
	message[Internal] = "Internal"                      // 500 grpc code 13 Internal 内部错误，底层错误 grpc错误
	message[Unavailable] = "Unavailable"                // 503 grpc code 14 Unavailable 服务不可用，稍后重试 grpc错误
	message[DataLoss] = "Data loss"                     // 500 grpc code 15 Data loss 数据损坏 非grpc错误
	message[Unauthenticated] = "Unauthenticated"        // 401 grpc code 16 Unauthenticated 未授权操作，鉴权中间件使用

	message[ServerCommonError] = "Server is error,please try later"
	message[RequestParamError] = "Request Params error"

	//custom error
	message[1001] = ""
	message[1002] = "用户名已经存在！"
	message[1003] = "Email已经存在！"
	message[1004] = "用户不存在，请先注册！"
	message[1005] = "用户未激活或已禁用，请联系管理员！"
	message[1006] = "密码不匹配！"
	message[1007] = "创建Token失败！"
	message[1008] = "Refresh Token不合法！"
	message[1009] = "密码须要使用数字、大小写字母、特殊字符中至少 2 种，且长度在 8-16 之间！"
	message[1010] = "用户名不可以包含特殊字符，且长度在 4-32 之间！"
	message[1011] = "昵称不可以包含特殊字符，长度在4-32之前！"
	message[1012] = "不合法的Email！"

}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _, ok := message[errCode]; ok {
		return true
	} else {
		return false
	}
}

func GetErrCode(grpcCode uint32) (uint32, error) {
	if errCode, ok := codeToErr[grpcCode]; ok {
		return errCode, nil
	} else {
		return 0, nil
	}
}
