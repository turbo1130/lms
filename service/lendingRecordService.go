package service

import (
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/model/request"
	"LibraryManagementSys/model/response"
	"LibraryManagementSys/utils"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
)

type LendingRecordService struct {
}

func (lr *LendingRecordService) GetLrListByPage(page int) (lrRes response.PageResult, err error) {
	var lrList []model.LendingRecord
	err = global.LMS_DB.Model(&model.LendingRecord{}).Where("deletedAt is NULL").Offset((page - 1) * 11).Limit(11).Order("jysj desc").Find(&lrList).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return lrRes, err
	}
	var total int64
	err = global.LMS_DB.Model(&model.LendingRecord{}).Where("deletedAt is NULL").Count(&total).Error
	if err != nil {
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return lrRes, err
	}
	lrRes.List = lrList
	lrRes.Page = page
	lrRes.PageSize = utils.CalPageSize(total)
	lrRes.Total = total
	return lrRes, nil
}

func (lr *LendingRecordService) GetLrListByCondition(lrCondition request.LendingRecCondition) (lrRes response.PageResult, err error) {
	var lrList []model.LendingRecord
	sqlDb := global.LMS_DB.Model(&model.LendingRecord{})
	err = lrListConditionSql(lrCondition, sqlDb).Scan(&lrList).Error
	if err != nil {
		return response.PageResult{}, err
	}
	var total int64
	sqlDb = global.LMS_DB.Model(&model.LendingRecord{})
	err = lrListConditionSql(lrCondition, sqlDb).Count(&total).Error
	lrRes.List = lrList
	lrRes.Page = lrCondition.Page
	lrRes.PageSize = utils.CalPageSize(total)
	lrRes.Total = total
	return lrRes, nil
}

func (lr *LendingRecordService) AddLrRec(lrRec model.LendingRecord) error {
	lrRec.UpdatedAt = time.Now()
	lrRec.CreatedAt = time.Now()
	// 查询是否可借出
	var bookCount int64
	global.LMS_DB.Model(&model.Book{}).Where("id = ? and stock > 0", lrRec.BookId).Count(&bookCount)
	if bookCount != 1 {
		return errors.New("该书库存不足，无法借阅！")
	}
	// 借出 开启事务
	lrRecMap := make(map[string]interface{})
	xtDb := global.LMS_DB.Begin()
	marshal, _ := json.Marshal(lrRec)
	_ = json.Unmarshal(marshal, &lrRecMap)
	delete(lrRecMap, "deleteAt")
	delete(lrRecMap, "hssj")
	delete(lrRecMap, "hsdjr")
	delete(lrRecMap, "hszhId")
	err := xtDb.Model(&model.LendingRecord{}).Create(&lrRecMap).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	columnsMap := make(map[string]interface{})
	columnsMap["jcs"] = gorm.Expr("jcs+1")
	columnsMap["stock"] = gorm.Expr("stock-1")
	err = xtDb.Model(&model.Book{}).Where("id = ?", lrRec.BookId).UpdateColumns(columnsMap).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	xtDb.Commit()
	return nil
}

func (lr *LendingRecordService) UpdateLrRec(lrRec model.LendingRecord) error {
	lrRec.UpdatedAt = time.Now()
	xtDb := global.LMS_DB.Begin()
	err := xtDb.Model(&model.LendingRecord{}).Where("id = ?", lrRec.ID).Updates(lrRec).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	xtDb.Commit()
	return nil
}

func (lr *LendingRecordService) UpdateLrRecHs(lrRec model.LendingRecord) error {
	lrRec.UpdatedAt = time.Now()
	// 还书
	xtDb := global.LMS_DB.Begin()
	err := xtDb.Model(&model.LendingRecord{}).Where("id = ?", lrRec.ID).Updates(lrRec).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	columnsMap := make(map[string]interface{})
	columnsMap["jcs"] = gorm.Expr("jcs-1")
	columnsMap["stock"] = gorm.Expr("stock+1")
	err = xtDb.Model(&model.Book{}).Where("id = ?", lrRec.BookId).UpdateColumns(columnsMap).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	xtDb.Commit()
	return nil
}

func (lr *LendingRecordService) DelLrRec(id string) error {
	if utils.IsEmpty(id) {
		return errors.New("借阅ID为空。")
	}
	xtDb := global.LMS_DB.Begin()
	err := xtDb.Model(&model.LendingRecord{}).Where("id = ?", id).UpdateColumn("deletedAt", time.Now()).Error
	if err != nil {
		xtDb.Rollback()
		global.LMS_LOG.Errorf("发生错误！Err: %s", err.Error())
		return err
	}
	xtDb.Commit()
	return nil
}

func lrListConditionSql(lrCondition request.LendingRecCondition, sqlDb *gorm.DB) *gorm.DB {
	if !utils.IsEmpty(lrCondition.Xh) {
		sqlDb.Where("xh = ?", lrCondition.Xh)
	}
	if !utils.IsEmpty(lrCondition.Jyr) {
		sqlDb.Where("jyr = ?", lrCondition.Jyr)
	}
	if !utils.IsEmpty(lrCondition.Jysj) {
		sqlDb.Where("jysj = ?", lrCondition.Jysj)
	}
	if !utils.IsEmpty(lrCondition.BookName) || !utils.IsEmpty(lrCondition.Isbn) {
		lrListConditionJoinSql(lrCondition, sqlDb)
	}
	return sqlDb
}

func lrListConditionJoinSql(lrCondition request.LendingRecCondition, sqlDb *gorm.DB) {
	sqlDb.Joins("left join books on lending_records.bookId = books.id")
	if !utils.IsEmpty(lrCondition.BookName) {
		sqlDb.Where("books.name = '" + lrCondition.BookName + "'")
	}
	if !utils.IsEmpty(lrCondition.Isbn) {
		sqlDb.Where("books.isbn = '" + lrCondition.Isbn + "'")
	}
}
