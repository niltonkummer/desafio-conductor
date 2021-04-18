package config

import "gorm.io/gorm"

type Config struct {
	HttpListen string
	DBDialect  gorm.Dialector
}
