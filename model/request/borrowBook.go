package request

import "time"

type BorrowBook struct {
	Id     string    `json:"id"`     // 借阅ID
	BookId string    `json:"bookId"` // 书ID
	Xh     string    `json:"xh"`     // 学号
	Jyr    string    `json:"jyr"`    // 借阅人
	Jydjr  string    `json:"jydjr"`  // 借阅登记人
	Jysj   time.Time `json:"jysj"`   // 借阅时间
}
