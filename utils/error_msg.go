package utils

import "sim-backend/models/common"

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_PASSWORD_WRONG   = 1001
	ERROR_TOKEN_EXIST      = 1002
	ERROR_TOKEN_RUNTIME    = 1003
	ERROR_TOKEN_WRONG      = 1004
	ERROR_TOKEN_TYPE_WRONG = 1005
	ERROR_USER_AUTHORITY   = 1006
	ERROR_USER_EXIST       = 1007
	ERROR_PASSWORD_MATCH = 1008
	ERROR_APPLICATION_EXIST = 1009

)

var errorMsgMap = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_TOKEN_EXIST:      "TOKEN不存在，请重新登录",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期，请重新登录",
	ERROR_TOKEN_WRONG:      "TOKEN不正确，请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误，请重新登录",
	ERROR_USER_AUTHORITY:   "该用户无权限",
	ERROR_USER_EXIST:       "用户不存在",
	ERROR_PASSWORD_MATCH:"原密码不正确",
	ERROR_APPLICATION_EXIST:"志愿信息未填报",
}

func GetErrMsg(code int) string {
	return errorMsgMap[code]
}

func Response(code int, data interface{}) common.Response {
	return common.Response{
		Status: code,
		Data:   data,
		Msg:    errorMsgMap[code],
	}
}

func ResponseWithError(code int, err error) common.Response{
	return common.Response{
		Status: code,
		Msg: errorMsgMap[code],
		Error:  err.Error(),
	}
}