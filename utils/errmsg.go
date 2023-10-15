package utils

const (
	SUCCESS      = 200
	REQUESTERROR = 400
	SERVERERROR  = 500

	ErrorUsernameUsed   = 1001
	ErrorPasswordWrong  = 1002
	ErrorUserNotExist   = 1003
	ErrorTokenNotExist  = 1004
	ErrorTokenRuntime   = 1005
	ErrorTokenWrong     = 1006
	ErrorTokenTypeWrong = 1007
)

var codeMsg = map[int]string{
	SUCCESS:             "OK",
	REQUESTERROR:        "请求格式错误",
	SERVERERROR:         "服务器处理错误",
	ErrorUsernameUsed:   "用户名已存在",
	ErrorPasswordWrong:  "密码错误",
	ErrorUserNotExist:   "用户不存在",
	ErrorTokenNotExist:  "TOKEN不存在",
	ErrorTokenRuntime:   "TOKEN已过期",
	ErrorTokenWrong:     "TOKEN不正确",
	ErrorTokenTypeWrong: "TOKEN格式不正确",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
