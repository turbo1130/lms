package model

import (
	"LibraryManagementSys/global"
	"time"
)

/*
   id: '12',
   book_id: '1',
   xh: '202310254',
   jyr: '张三',
   jysj:'2023-01-20 10:09:53',
   hssj:'-',
   name: '活着',
   author: '余华',
   isbn: 'SNB1234567',
   publisher: '人民出版社',
   jydjr:'张三',
   hsdjr:'-',
   isH: false
*/

type LendingRecord struct {
	global.LmsModel
	BookId string    `gorm:"column:bookId;comment:书ID;size:32" json:"bookId"`
	JyzhId string    `gorm:"column:jyzhId;comment:借阅操作账户ID;size:32" json:"jyzhId"`
	HszhId string    `gorm:"column:hszhId;comment:还书操作账户ID;size:32" json:"hszhId"`
	Xh     string    `gorm:"column:xh;comment:学号" json:"xh"`
	Jyr    string    `gorm:"column:jyr;comment:借阅人" json:"jyr"`
	Jydjr  string    `gorm:"column:jydjr;comment:借阅登记人" json:"jydjr"`
	Hsdjr  string    `gorm:"column:hsdjr;comment:还书登记人" json:"hsdjr"`
	Jysj   time.Time `gorm:"column:jysj;comment:借出时间" json:"jysj"`
	Hssj   time.Time `gorm:"column:hssj;comment:还书时间" json:"hssj"`
	IsH    bool      `gorm:"column:isH;comment:是否还书，已还true，未还false" json:"isH"`
}
