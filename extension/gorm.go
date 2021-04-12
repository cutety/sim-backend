package extension

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sim-backend/utils/logger"
)

const (
	MySQLError_DuplicateEntry = 1062
)

func IsMySQLDuplicateEntryError(err error) bool {
	if err == nil {
		return false
	}

	sqlErr, ok := err.(*mysql.MySQLError)
	// 唯一索引冲突错误
	return ok && sqlErr.Number == MySQLError_DuplicateEntry
}

var DB *gorm.DB


func InitDB() {
	logger.Info("MySQL starting...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("extensions.mysql.username"),
		viper.GetString("extensions.mysql.password"),
		viper.GetString("extensions.mysql.host"),
		viper.GetString("extensions.mysql.port"),
		viper.GetString("extensions.mysql.db"),
	)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db.SingularTable(false)
	db.LogMode(true)
	//db.SetLogger(logger.NewGormLogger())
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)

	DB = db
	logger.Info("MySQL is successfully connected")
}

