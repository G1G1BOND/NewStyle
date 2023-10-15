package message

const (
	Success      = 200
	Error        = 500
	InvalidParam = 400

	InvalidEmail           = 10001
	RepeatSending          = 10002
	WrongCode              = 10003
	RepeatEmail            = 10004
	RepeatName             = 10005
	WrongPasswordFormat    = 10006
	WrongAccountOrPassword = 10007
	IconTooBig             = 10009
	WrongPictureFormat     = 10010
	NilNickName            = 10011
	NilName                = 10013

	NotFoundName = 30002
)

var msg map[int]string

func init() {
	msg = make(map[int]string)
	msg[Success] = "ok"
	msg[Error] = "服务器内部错误"

	msg[InvalidEmail] = "邮箱格式不正确"
	msg[RepeatSending] = "发送过于频繁"
	msg[InvalidParam] = "参数解析异常"
	msg[WrongCode] = "验证码错误"
	msg[RepeatEmail] = "该邮箱已经注册"
	msg[RepeatName] = "该姓名已被使用"
	msg[NilName] = "用户名不能为空"
	msg[WrongAccountOrPassword] = "账号或密码不正确"
	msg[WrongPasswordFormat] = "密码格式不正确"
	//msg[UserNotLogin] = "用户未登录"
	msg[IconTooBig] = "只允许2MB以下的图片作为头像"
	msg[WrongPictureFormat] = "不支持该格式的图片"
	msg[NilNickName] = "昵称不能为空"

	msg[NotFoundName] = "找不到该用户名"
}

func GetMsg(code int) string {
	return msg[code]
}
