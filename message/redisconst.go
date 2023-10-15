package message

import "time"

const (
	VerificationCodeKey    = "newStyle:user:code:"
	VerificationCodeKeyTTL = time.Second * 60 * 3

	UserLoginInfo    = "newStyle:user:token:"
	UserLoginInfoTTL = time.Hour * 24 * 30
)
