package db

import (
	"log"

	"github-scanner/models"
	"github-scanner/utils"

	"gorm.io/gorm"
)

func SaveScan(db *gorm.DB, scanWrapper models.ScanWrapper) error {
	scan, vulnDetails, riskFactors, severityCounts, scanningRules, excludedPaths := utils.ConvertScanWrapperToDB(scanWrapper)

	tx := db.Begin()

	if err := tx.Create(&scan).Error; err != nil {
		tx.Rollback()
		log.Printf(utils.DB_SCAN_RULES_INSERT_ERROR, err)
		return err
	}

	if len(vulnDetails) > 0 {
		if err := tx.Model(&models.VulnerabilityDetails{}).Create(&vulnDetails).Error; err != nil {
			tx.Rollback()
			log.Printf(utils.DB_VULN_INSERT_ERROR, err)
			return err
		}
	}

	if len(riskFactors) > 0 {
		if err := tx.Model(&models.RiskFactor{}).Create(&riskFactors).Error; err != nil {
			tx.Rollback()
			log.Printf(utils.DB_RISK_FACTOR_INSERT_ERROR, err)
			return err
		}
	}

	if len(severityCounts) > 0 {
		if err := tx.Model(&models.SeverityCount{}).Create(&severityCounts).Error; err != nil {
			tx.Rollback()
			log.Printf(utils.DB_SEV_COUNT_INSERT_ERROR, err)
			return err
		}
	}

	if len(scanningRules) > 0 {
		if err := tx.Model(&models.ScanningRule{}).Create(&scanningRules).Error; err != nil {
			tx.Rollback()
			log.Printf(utils.DB_SCAN_RULES_INSERT_ERROR, err)
			return err
		}
	}

	if len(excludedPaths) > 0 {
		if err := tx.Model(&models.ExcludedPath{}).Create(&excludedPaths).Error; err != nil {
			tx.Rollback()
			log.Printf(utils.DB_EXC_PATH_INSERT_ERROR, err)
			return err
		}
	}

	tx.Commit()
	log.Println(utils.DB_SUCCESS_SAVE)
	return nil
}
