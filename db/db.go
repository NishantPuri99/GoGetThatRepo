package db

import (
	"github-scanner/models"
	"github-scanner/utils"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(sqlite.Open("scanner.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(utils.DB_CONNECT_FAIL, err)
		return nil, err
	}
	err = db.AutoMigrate(&models.Scan{}, &models.VulnerabilityDetails{}, &models.RiskFactor{}, &models.SeverityCount{}, &models.ScanningRule{}, &models.ExcludedPath{})
	if err != nil {
		log.Fatal(utils.DB_MIGRATE_FAIL, err)
		return nil, err
	}

	log.Println(utils.DB_TABLE_CREATION_SUCCESS)
	return db, nil
}
