package initialize

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	"LibraryManagementSys/model"
	"LibraryManagementSys/utils"
	"gorm.io/gorm"
	"os"
	"strings"
)

func DbInit() *gorm.DB {
	dbType := global.LMS_CONFIG.System.DbType
	// 后续可扩展支持其他数据库
	switch dbType {
	case "mysql":
		return MysqlInit()
	default:
		return MysqlInit()
	}
}

// RegisterTables
// 创建数据库表
func RegisterTables(db *gorm.DB) {
	tx := db.Begin() // 开启事务
	err := tx.AutoMigrate(
		&model.Book{},
		&model.LendingRecord{},
		&model.Role{},
		&model.User{},
		&model.Version{},
	)
	if err != nil {
		tx.Rollback() // 回滚
		global.LMS_LOG.Errorf("创建数据库表失败！")
		os.Exit(0)
	}
	tx.Commit() // 提交事务
}

// DataInit
// 根据版本号进行数据初始化
func DataInit(db *gorm.DB) {
	tx := db.Begin()
	version := global.LMS_CONFIG.System.Version
	path := constant.SqlInitPath + version
	// 查找version表是否有该版本信息，存在版本记录则不进行该版本数据初始化
	// tips: 这里逻辑还可以修改为跨版本升级，若当前版本为v3.0，而上线版本还是v1.0，在升级v3.0时，
	// 		 可根据上线版本的version表中最新记录v1.0依次循环执行sql_init下的v2.0目录脚本和v3.0目录脚本
	// 		 达到升级到v3.0（本程序逻辑直接默认没有跨版本升级的情况）
	var resCount int64
	tx.Model(&model.Version{}).Where("version = ?", version).Count(&resCount)
	if resCount != 0 {
		tx.Commit()
		global.LMS_LOG.Infof("已是最新版本，版本号为：%s，无需更新数据!", version)
		return
	}
	// 初始化数据
	global.LMS_LOG.Infof("正在初始化%s版本的数据...", version)
	files, _ := os.ReadDir(path)
	for _, file := range files {
		if strings.HasSuffix(file.Name(), constant.SqlFileSuffix) {
			global.LMS_LOG.Infof("正在执行%s文件...", file.Name())
			contents, _ := os.ReadFile(path + constant.FileFix + file.Name())
			res := strings.Split(string(contents), constant.SqlSplit)
			for _, sql := range res {
				if !utils.IsBlankString(sql) {
					tx.Exec(sql)
				}
			}
			global.LMS_LOG.Infof("%s文件执行完成...", file.Name())
		}
	}
	// 插入版本信息
	global.LMS_LOG.Infof("正在记录%s版本信息...", version)
	verObj := model.Version{Version: version}
	verObj.ID = utils.GetRandomId()
	tx.Create(&verObj)
	global.LMS_LOG.Infof("记录%s版本信息完成...", version)

	tx.Commit()
	global.LMS_LOG.Infof("%s版本的数据初始化完成...", version)
}
