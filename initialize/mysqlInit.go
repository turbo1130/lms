package initialize

import (
	"LibraryManagementSys/core"
	"LibraryManagementSys/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlInit
// mysql数据库初始化
func MysqlInit() *gorm.DB {
	m := global.LMS_CONFIG.Mysql
	if m.DbName == "" {
		return nil
	}
	db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{Logger: core.NewGormZapLogger()})
	if err != nil {
		global.LMS_LOG.Errorf("数据库连接失败：%s", err)
	} else {
		global.LMS_LOG.Infof("数据库连接成功，连接的数据名为：%s", global.LMS_CONFIG.Mysql.DbName)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(m.MaxIdleConns)
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	return db
}
