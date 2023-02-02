package config

type System struct {
	Name             string `mapstructure:"name" json:"name" yaml:"name"`
	Version          string `mapstructure:"version" json:"version" yaml:"version"`
	DbType           string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
	RouterPrefix     string `mapstructure:"router-prefix" json:"routerPrefix" yaml:"router-prefix"`
	TransferPwDecode bool   `mapstructure:"transfer-pw-decode" json:"transferPwDecode" yaml:"transfer-pw-decode"`
}
