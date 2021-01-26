package error_msg

const (
	SUCCESS = 200
	ERROR   = 500

	ERROR_PASSWORD_WRONG   = 1001
	ERROR_TOKEN_EXIST      = 1002
	ERROR_TOKEN_RUNTIME    = 1003
	ERROR_TOKEN_WRONG      = 1004
	ERROR_TOKEN_TYPE_WRONG = 1005
	ERROR_USER_AUTHORITY   = 1006
)

var errorMsgMap = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_TOKEN_EXIST:      "TOKEN已存在，请重新登录",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期，请重新登录",
	ERROR_TOKEN_WRONG:      "TOKEN不正确，请重新登录",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误，请重新登录",
	ERROR_USER_AUTHORITY:   "该用户无权限",
}

func GetErrMsg(code int) string {
	return errorMsgMap[code]
}
