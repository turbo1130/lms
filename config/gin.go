package config

type Gin struct {
	Port string `mapstructure:"port" json:"port" yaml:"port"`
}
