package setup

import (
	"gorm.io/gorm"
)

func SetupDB(dialect gorm.Dialector) (*gorm.DB, error) {
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
