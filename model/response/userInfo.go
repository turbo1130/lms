package response

type UserInfo struct {
	Id       string `json:"id"`                              // ID
	Username string `json:"username"`                        // 登录名
	Phone    string `json:"phone"`                           // 电话
	Email    string `json:"email"`                           // 邮箱
	NickName string `gorm:"column:nickName" json:"nickName"` // 昵称
}
