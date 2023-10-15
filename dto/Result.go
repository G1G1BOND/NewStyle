package dto

import "go-v1/message"

type Result struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"message"`
}

func Fail(code int, err error) *Result {
	result := &Result{
		Code: code,
		Msg:  message.GetMsg(code),
	}
	if err != nil {
		result.Data = err.Error()
	}
	return result
}

func Success(code int, data interface{}) *Result {
	return &Result{
		Code: code,
		Msg:  message.GetMsg(code),
		Data: data,
	}
}
