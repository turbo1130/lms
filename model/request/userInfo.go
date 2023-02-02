package request

import "time"

type UserInfo struct {
	Id         string    `json:"id"`         // ID
	Phone      string    `json:"phone"`      // 电话
	Email      string    `json:"email"`      // 邮箱
	NickName   string    `json:"nickName"`   // 昵称
	Password   string    `json:"password"`   // 密码
	ChangeTime time.Time `json:"changeTime"` // 更改时间
}
