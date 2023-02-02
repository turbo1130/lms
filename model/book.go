package model

import (
	"LibraryManagementSys/global"
)

type Book struct {
	global.LmsModel
	Num       string `gorm:"column:num;comment:书号" json:"num"`              // 书号
	Name      string `gorm:"column:name;comment:书名" json:"name"`            // 书名
	Author    string `gorm:"column:author;comment:作者" json:"author"`        // 作者
	Isbn      string `gorm:"column:isbn;comment:ISBN" json:"isbn"`          // ISBN
	Publisher string `gorm:"column:publisher;comment:出版社" json:"publisher"` // 出版社
	Price     string `gorm:"column:price;comment:价格" json:"price"`          // 价格
	Jcs       uint   `gorm:"column:jcs;comment:借出数" json:"jcs"`             // 借出数
	Stock     uint   `gorm:"column:stock;comment:库存" json:"stock"`          // 库存
	IsJcs     bool   `gorm:"-" json:"-"`                                    // 需要更新jcs为0时，需设置为true才能正确更新
	IsStock   bool   `gorm:"-" json:"-"`                                    // 需要更新stock为0时，需设置为true才能正确更新
}

/*
   id: '12',
   num: '12-4561',
   name: '活着',
   author: '余华',
   isbn: 'SNB1234567',
   publisher: '人民出版社',
   price: '12.50',
   jcs: 12,
   stock: 30
*/
