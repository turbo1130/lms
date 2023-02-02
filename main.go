package main

import (
	"LibraryManagementSys/core"
	"LibraryManagementSys/global"
	"LibraryManagementSys/initialize"
)

// @title LMS
// @version 1.0
// @description 图书借阅系统接口
// @contact.name VinceZ
// @contact.url https://gitee.com/turbo30/library-management-sys_go
// @contact.email turbochang@foxmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /
func main() {
	core.Viper() // 启动环境在 constant/configConstant.go 中ENV配置

	core.LmsInitLogger() // 日志初始化

	global.LMS_DB = initialize.DbInit() // 连接数据库
	if global.LMS_DB != nil {
		initialize.RegisterTables(global.LMS_DB) // 创建表
		initialize.DataInit(global.LMS_DB)       // 初始化数据
		// 程序结束前关闭数据库链接
		db, _ := global.LMS_DB.DB()
		defer db.Close()
	}

	initialize.RedisInit() // 连接redis
	defer global.LMS_REDIS.Close()

	core.RunGinServer() // 启动服务
}
