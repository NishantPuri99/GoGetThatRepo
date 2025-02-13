package db

import (
	"github-scanner/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("scanner.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Scan{}, &models.VulnerabilityDetails{}, &models.RiskFactor{}, &models.SeverityCount{}, &models.ScanningRule{}, &models.ExcludedPath{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
		return nil, err
	}

	log.Println("Database and tables created successfully.")
	return db, nil
}
