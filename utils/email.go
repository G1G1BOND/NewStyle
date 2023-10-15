package utils

import (
	"fmt"
	"go-v1/config"
	gomail "gopkg.in/mail.v2"
)

func SendCode(code string, e string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.FROM)
	m.SetHeader("To", e)
	//m.SetHeader("To", e)
	m.SetHeader("Subject", "验证码")
	m.SetBody("text/plain", "您的验证码为："+code+" （一分钟内有效）")
	d := gomail.NewDialer("smtp.qq.com", 465, config.FROM, config.PASSWORD)
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("发送结束")
}
