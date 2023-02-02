package request

import "time"

type ReturnBook struct {
	Id     string    `json:"id"`     // 借阅ID
	BookId string    `json:"bookId"` // 还书ID
	Hsdjr  string    `json:"hsdjr"`  // 还书登记人
	Hssj   time.Time `json:"hssj"`   // 还书时间
}
