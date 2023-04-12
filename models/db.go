package models

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strconv"
	"time"
	"websocket/config"
)

var DB *gorm.DB

func ConnectDB() {
	conf := config.GetConfig()
	log.Debug("Set meta db driver " + conf.Server.DbType)

	//var db *gorm.DB
	var err error

	newLogger := logger.New(
		log.StandardLogger(),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Silent,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	confGorm := gorm.Config{Logger: newLogger}

	log.Info("Init DB " + conf.Server.DbType + " " + strconv.Itoa(int(log.GetLevel())))

	switch conf.Server.DbType {
	case "sqlite":
		DB, err = gorm.Open(sqlite.Open(conf.Db.DbDsn), &confGorm)
	case "mysql":
		DB, err = gorm.Open(sqlite.Open(conf.Db.DbDsn), &confGorm)
	case "postgres":
		DB, err = gorm.Open(sqlite.Open(conf.Db.DbDsn), &confGorm)
	}

	if err != nil {
		panic("failed to connect database " + err.Error())
	}

	if log.GetLevel() >= log.DebugLevel {
		DB.Debug()
	}

	err = DB.Debug().AutoMigrate(
		UserConnection{},
	)

	if err != nil {
		log.Panic(err)
	}
}
