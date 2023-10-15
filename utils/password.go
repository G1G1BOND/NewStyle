package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密
func Encrypt(password string) (string, error) {
	// bcrypt.DefaultCost=10 可以输入数字 范围4-31
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 失败时返回 error
	if err != nil {
		return "", err
	}
	// 成功时返回 hash
	return string(hashPassword), nil
}

// 密码解密
func Decode(hassPassword, password string) bool {
	// 失败时返回 false
	if err := bcrypt.CompareHashAndPassword([]byte(hassPassword), []byte(password)); err != nil {
		return false
	}
	// 成功时返回 true
	return true
}
