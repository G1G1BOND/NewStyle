package dto

import (
	"go-v1/model"
	"time"
)

type UserDto struct {
	Email string `json:"email"`
	//NickName string `json:"nickName"`
	Avatar string `json:"avatar"`
	Token  string `json:"token,omitempty"`
}

type MomentDto struct {
	Name      string    `json:"name"`
	Avatar    string    `json:"avatar"`
	Picture   string    `json:"picture"`
	Content   string    `json:"content"`
	Like      string    `json:"like"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func BuildUser(user *model.User, token string) *UserDto {
	return &UserDto{
		Email: user.Email,
		//NickName: user.NickName,
		Token:  token,
		Avatar: user.Avatar,
	}
}
