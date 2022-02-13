package middleware

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zhihu.com/m/config"
	"zhihu.com/m/define"
)

var (
	WriteDB *gorm.DB
	ReadDB  *gorm.DB

	writeMaxOpenConns    = 8192
	writeConnMaxIdleTime = 10 * time.Minute

	readMaxOpenConns    = 8192
	readConnMaxIdleTime = 10 * time.Minute
)

func getWriteDB(dbConfig *define.ConfigDB, gormConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.ReadUsername, dbConfig.ReadPassword, dbConfig.ReadHost, dbConfig.Database, dbConfig.Charset)

	if gormConfig == nil {
		gormConfig = &gorm.Config{}
	}

	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func getReadDB(dbConfig *define.ConfigDB, gormConfig *gorm.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=Local", dbConfig.ReadUsername, dbConfig.ReadPassword, dbConfig.ReadHost, dbConfig.Database, dbConfig.Charset)

	if gormConfig == nil {
		gormConfig = &gorm.Config{}
	}

	return gorm.Open(mysql.Open(dsn), gormConfig)
}

func openWriteConn(dbConfig *define.ConfigDB) *gorm.DB {
	db, err := getWriteDB(dbConfig, nil)
	if err != nil {
		panic(fmt.Errorf("mysql init is err: %v", err))
	}

	realDB, err := db.DB()

	if err != nil {
		panic(fmt.Errorf("get mysql DB is err:%v", err))
	}

	realDB.SetMaxIdleConns(writeMaxOpenConns / 2)
	realDB.SetMaxOpenConns(writeMaxOpenConns)
	realDB.SetConnMaxLifetime(writeConnMaxIdleTime)

	return db
}

func openReadConn(dbConfig *define.ConfigDB) *gorm.DB {
	db, err := getReadDB(dbConfig, nil)
	if err != nil {
		panic(fmt.Errorf("mysql init is err: %v", err))
	}

	realDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("get mysql DB is err: %v", err))
	}

	realDB.SetMaxIdleConns(readMaxOpenConns / 2)
	realDB.SetMaxOpenConns(readMaxOpenConns)
	realDB.SetConnMaxLifetime(readConnMaxIdleTime)

	return db
}

// InitMysqlConn ...
func InitMysqlConn() {
	config, err := config.GetDBConfig("config/mysql")

	if err != nil {
		panic(err)
	}

	WriteDB = openWriteConn(config)
	ReadDB = openReadConn(config)
}
