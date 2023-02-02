package service

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"errors"
	"time"
)

type BookService struct {
}

func (b *BookService) GetBookListByPage(page int) (pageRes response.PageResult, err error) {
	var bookList []model.Book
	err = global.LMS_DB.Where("deletedAt is NULL").Offset((page - 1) * 11).Limit(11).Order("jcs desc, stock").Find(&bookList).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return pageRes, err
	}
	var total int64
	global.LMS_DB.Model(&model.Book{}).Where("deletedAt is NULL or deletedAt != 1").Count(&total)
	pageRes.List = bookList
	pageRes.PageSize = utils.CalPageSize(total)
	pageRes.Page = page
	pageRes.Total = total
	return pageRes, nil
}

func (b *BookService) GetBookListByConditions(p request.PageReq) (response.PageResult, error) {
	book := p.Data.(model.Book)
	var bookList []model.Book
	err := global.LMS_DB.Where(&book).Offset((p.Page - 1) * 11).Limit(11).Order("jcs desc, stock").Find(&bookList).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return response.PageResult{}, err
	}
	var total int64
	global.LMS_DB.Model(&model.Book{}).Where(&book).Count(&total)
	var pageRes response.PageResult
	pageRes.List = bookList
	pageRes.PageSize = utils.CalPageSize(total)
	pageRes.Page = p.Page
	pageRes.Total = total
	return pageRes, nil
}

func (b *BookService) GetPublisherList() (pageRes response.PageResult, err error) {
	var publishers []string
	err = global.LMS_DB.Model(&model.Book{}).Select("publisher").Distinct("publisher").Find(&publishers).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return response.PageResult{}, err
	}
	pageRes.List = publishers
	return pageRes, nil
}

func (b *BookService) AddBook(book model.Book) error {
	book.ID = utils.GetRandomId()
	book.UpdatedAt = time.Now()
	book.CreatedAt = time.Now()
	err := global.LMS_DB.Create(&book).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	return nil
}

func (b *BookService) UpdateBook(book model.Book) error {
	bookDao := book
	bookDao.ID = constant.EMPTY
	bookDao.UpdatedAt = time.Now()
	db := global.LMS_DB.Begin()
	err := db.Model(&model.Book{}).Where("id = ?", book.ID).Updates(bookDao).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		db.Rollback()
		return err
	}
	// 解决值为0的更新操作
	upMap := make(map[string]interface{})
	if book.IsJcs && book.Jcs == 0 {
		upMap["jcs"] = 0
	}
	if book.IsStock && book.Stock == 0 {
		upMap["stock"] = 0
	}
	err = db.Model(&model.Book{}).Where("id = ?", book.ID).Updates(upMap).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func (b *BookService) DelBook(book model.Book) error {
	var count int64
	// 查询该书是否存在借出记录
	err := global.LMS_DB.Model(&model.Book{}).Where("id = ? and jcs = 0", book.ID).Count(&count).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	// 如果该书没有借出，则可以删除(软删除)
	if count == 1 {
		err = global.LMS_DB.Model(&model.Book{}).Delete(&book).Error
		if err != nil {
			global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
			return err
		}
	} else {
		return errors.New("该书借出数不为0，不能进行删除操作！")
	}
	return nil
}
