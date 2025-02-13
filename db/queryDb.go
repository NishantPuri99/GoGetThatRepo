package db

import (
	"github-scanner/models"

	"gorm.io/gorm"
)

func QueryVulns(db *gorm.DB, severity string) ([]models.Vulnerability, error) {
	var vulnerabilities []models.VulnerabilityDetails

	err := db.Where("severity = ?", severity).Find(&vulnerabilities).Error
	if err != nil {
		return nil, err
	}
	var responses []models.Vulnerability

	for _, vuln := range vulnerabilities {
		var riskFactors []string
		db.Model(&models.RiskFactor{}).
			Where("vuln_id = ?", vuln.VulnID).
			Pluck("risk_factor", &riskFactors)

		responses = append(responses, models.Vulnerability{
			ID:             vuln.VulnID,
			Severity:       vuln.Severity,
			CVSS:           vuln.CVSS,
			Status:         vuln.Status,
			PackageName:    vuln.PackageName,
			CurrentVersion: vuln.CurrentVersion,
			FixedVersion:   vuln.FixedVersion,
			Description:    vuln.Description,
			PublishedDate:  vuln.PublishedDate,
			Link:           vuln.Link,
			RiskFactors:    riskFactors,
		})
	}

	return responses, nil
}
