package models

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Setup() {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: 500 * time.Millisecond, // Slow SQL threshold
			LogLevel:      logger.Info,            // Log level
			Colorful:      true,                   // Disable color
		},
	)

	DB, err = gorm.Open(sqlite.Open("task.db"),
		&gorm.Config{
			Logger: newLogger,
		})

	if err != nil {
		log.Fatal(err.Error())
	}

}

func GetDB() *gorm.DB {
	return DB
}
