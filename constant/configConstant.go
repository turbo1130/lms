package constant

const (
	// 环境启动读取优先级：传参 > 启动参数 > 环境变量 > Env配置
	ConfigEnv         = "LMS_CONFIG"
	ENV               = "dev"
	DEV               = "dev"
	RELEASE           = "release"
	DEFAULT           = "default"
	ConfigDefaultFile = "config-default.yaml"
	ConfigDevFile     = "config-dev.yaml"
	ConfigReleaseFile = "config-release.yaml"

	// 日志
	ZapLogFormatConsole = "console"
	ZapLogFormatJson    = "json"

	// 数据库初始化
	SqlInitPath = "./source/sql_init/"
)
