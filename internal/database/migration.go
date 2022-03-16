package database

import (
	"github.com/PatriciaChebet/rest_api_with_go/internal/comment"
	"github.com/jinzhu/gorm"
)

//MigrateDB - migrates our database and creates our comments table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
