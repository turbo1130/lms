package config

type Zap struct {
	Level        string `mapstructure:"level" json:"level" yaml:"level"`
	GormLevel    string `mapstructure:"gorm-level" json:"gormLevel" yaml:"gorm-level"`
	Format       string `mapstructure:"format" json:"format" yaml:"format"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Prefix       string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	LogInConsole bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
