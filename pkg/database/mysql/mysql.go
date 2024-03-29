package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"includemy/pkg/config"
	"log"
)

func ConnectDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.LoadDataSourceName()), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
