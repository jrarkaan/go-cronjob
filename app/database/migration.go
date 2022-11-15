package database

import (
	"github.com/jrarkaan/go-cronjob/dto"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&dto.MProduct{})
}
