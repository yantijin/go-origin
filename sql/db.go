package sql

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Open(dsn string, config *gorm.Config, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if db, err = gorm.Open(mysql.Open(dsn), config); err != nil {
		log.Errorf("Open database failed: %s", err.Error())
	}

	if err = db.AutoMigrate(models...); err != nil {
		log.Errorf("Autommigration database error: %s", err.Error())
	}

	return
}

func Db() *gorm.DB {
	return db
}
