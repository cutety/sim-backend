package extension

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"sim-backend/utils/logger"
)

var db *gorm.DB

func init() {
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
	db.SingularTable(true)
	db.LogMode(true)
	db.SetLogger(logger.NewGormLogger())
	db.DB().SetMaxIdleConns(50)
	db.DB().SetMaxOpenConns(200)

	logger.Info("MySQL is sucessfully connected")
}

func GetDb() *gorm.DB {
	return db
}
