package core

import (
	"LibraryManagementSys/constant"
	"LibraryManagementSys/global"
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

// Viper
// 环境启动读取优先级：传参 > 命令行 > 环境变量 > Env配置
func Viper(path ...string) {
	var config string

	if len(path) == 0 {
		// 读取并解析命令行 -conf 参数，赋值给变量 config
		flag.StringVar(&config, "conf", "", "choose config file.")
		flag.Parse()
		if config == "" { // 判断命令行参数是否为空
			if configEnv := os.Getenv(constant.ConfigEnv); configEnv != "" { // 获取环境变量
				SetConfig(&config, configEnv)
			} else { // 无环境变量根据constant.Env获取环境启动方式
				SetConfig(&config, constant.ENV)
			}
		} else { // 命令行参数不为空 将值赋值于config
			fmt.Printf("您正在使用命令行的-conf参数传递的值,config的路径为%s\n", config)
		}
	} else { // 函数传递的可变参数的第一个值赋值于config（yaml配置文件路径）
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%s\n", config)
	}

	v := viper.New()
	v.SetConfigFile(config) // 配置文件路径
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监听配置文件修改，热更新
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.LMS_CONFIG); err != nil {
			fmt.Println(err)
		}
	})

	// 解析配置文件映射到结构体中
	if err = v.Unmarshal(&global.LMS_CONFIG); err != nil {
		fmt.Println(err)
	}

	// 设置到全局变量中
	global.LMS_VP = v
}

func SetConfig(config *string, configEnv string) {
	switch configEnv {
	case constant.DEV:
		*config = constant.ConfigDevFile
		gin.SetMode(gin.DebugMode)
		fmt.Printf("您正在使用%s环境,gin模式为%s环境,config的路径为%s\n", constant.DEV, gin.DebugMode, constant.ConfigDevFile)
	case constant.RELEASE:
		*config = constant.ConfigReleaseFile
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("您正在使用%s环境,gin模式为%s环境,config的路径为%s\n", constant.RELEASE, gin.ReleaseMode, constant.ConfigReleaseFile)
	default:
		*config = constant.ConfigDefaultFile
		gin.SetMode(gin.TestMode)
		fmt.Printf("您正在使用%s环境,gin模式为%s环境,config的路径为%s\n", constant.DEFAULT, gin.TestMode, constant.ConfigDefaultFile)
	}
}
