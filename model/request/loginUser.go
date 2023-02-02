package request

import "time"

type LoginUser struct {
	Username  string    `json:"username"`  // 用户名
	Password  string    `json:"password"`  // 密码
	LoginTime time.Time `json:"loginTime"` // 登陆时间
}
