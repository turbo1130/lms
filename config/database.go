package config

// DsnProvider
// 根据不同的数据库生成对应的连接地址
type DsnProvider interface {
	Dsn() string
}

type GeneralDB struct {
	Ip           string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	DbName       string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Charset      string `mapstructure:"charset" json:"charset" yaml:"charset"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      string `mapstructure:"logMode" json:"logMode" yaml:"log-mode"`
	LogZap       bool   `mapstructure:"logZap" json:"logZap" yaml:"log-zap"`
}
