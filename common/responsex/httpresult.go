package responsex

import (
	"bodhiadmin/common/errorx"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
)

// http返回
func HttpResult(req *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		//requestUuid := insertAPIRecord(recordRpc, "200", req)
		//成功返回
		r := Success("OK", resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := errorx.ServerCommonError
		errmsg := "server error, try again later"

		causeErr := errors.Cause(err)                  // err类型
		if e, ok := causeErr.(*errorx.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
			logx.WithContext(req.Context()).Errorf("[API-ERR] %+v ", errmsg)
		} else {
			// 将GRPC错误CODE 映射为 HTTP code
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())

				if errorx.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = errorx.MapErrMsg(errcode)
				} else {
					// 转为对应的http标准错误码
					errcode, _ = errorx.GetErrCode(grpcCode)
					if errcode == 0 {
						errcode = errorx.ServerCommonError
					}
					errmsg = gstatus.Message()
				}
				logx.WithContext(req.Context()).Errorf("[RPC-ERR] %+v ", gstatus.Message())
			}
		}

		if int(errcode) < 1000 {
			httpx.WriteJson(w, int(errcode), Error(errcode, errmsg))
		} else {
			//自定义模块错误都返回200，以自定义错误码返回错误信息
			httpx.WriteJson(w, http.StatusOK, Error(errcode, errmsg))
		}
	}
}
