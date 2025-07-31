package db

import (
	"fmt"
	"log"

	"github.com/Aritiaya50217/CodingTestByTriofarm/internal/domain"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MSSQL: %w", err)
	}

	// Auto Migration
	err = db.AutoMigrate(&domain.Topic{})
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	// Check if default topic exists; if not, create one
	var count int64
	db.Model(&domain.Topic{}).Count(&count)

	if count == 0 {
		defaultTopic := []domain.Topic{
			{Name: domain.Medicine},
			{Name: domain.Vitamins},
			{Name: domain.Microorganisms},
			{Name: domain.Brands},
		}
		if err := db.Create(&defaultTopic).Error; err != nil {
			log.Printf("failed to create default topic: %v", err)
		} else {
			log.Println("default topic created")
		}
	}

	return db, nil
}
