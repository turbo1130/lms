package config

type Server struct {
	System System `mapstructure:"system" json:"system" yaml:"system"`
	JWT    JWT    `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap    Zap    `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis  Redis  `mapstructure:"redis" json:"redis" yaml:"redis"`
	// gorm
	Mysql Mysql `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Excel Excel `mapstructure:"excel" json:"excel" yaml:"excel"`
	// gin
	Gin Gin `mapstructure:"gin" json:"gin" yaml:"gin"`
	// 跨域配置
	// Cors CORS `mapstructure:"cors" json:"cors" yaml:"cors"`
}
