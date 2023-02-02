package config

import (
	"fmt"
)

type Mysql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"` // mapstructure:",squash" 将该结构体的字段提到父结构中
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true&loc=Local", m.Username, m.Password, m.Ip, m.Port, m.DbName, m.Charset)
}

func (m *Mysql) GetLogMode() string {
	return m.LogMode
}
