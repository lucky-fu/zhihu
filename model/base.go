package model

import (
	"gorm.io/gorm"
	"zhihu.com/m/middleware"
)

// GetWriteDB ...
func GetWriteDB() *gorm.DB {
	return middleware.WriteDB
}

// GetReadDB ...
func GetReadDB() *gorm.DB {
	return middleware.ReadDB
}
