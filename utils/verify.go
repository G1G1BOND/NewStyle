package utils

import "regexp"

func VerifyEmailFormat(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//func VerifyMobileFormat(mobileNum string) bool {
//	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"
//
//	reg := regexp.MustCompile(regular)
//	return reg.MatchString(mobileNum)
//}

func VerifyPasswordFormat(password string) bool {
	var (
		hasNum     = false
		hasChar    = false
		hasSpecial = false
	)
	//检查长度
	if len(password) < 8 || len(password) > 16 {
		return false
	}
	for _, c := range password {
		switch {
		case c >= '0' && c <= '9':
			hasNum = true
		case (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'):
			hasChar = true
		default:
			hasSpecial = true
		}
	}
	return hasChar && hasNum && !hasSpecial
}
