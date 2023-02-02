package request

type LendingRecCondition struct {
	Page     int    `json:"page"`     // 页码
	Xh       string `json:"xh"`       // 学号
	Jyr      string `json:"jyr"`      // 借阅人
	Jysj     string `json:"jysj"`     // 借阅时间
	BookName string `json:"bookName"` // 书名
	Isbn     string `json:"isbn"`     // isbn
}
