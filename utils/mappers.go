package utils

import (
	"github-scanner/models"
)

func ConvertScanWrapperToDB(scanWrapper models.ScanWrapper) (
	models.Scan,
	[]models.VulnerabilityDetails,
	[]models.RiskFactor,
	[]models.SeverityCount,
	[]models.ScanningRule,
	[]models.ExcludedPath) {

	scanResults := scanWrapper.ScanResult

	scan := models.Scan{
		ScanID:          scanResults.ScanID,
		Timestamp:       scanResults.Timestamp,
		ScanStatus:      scanResults.ScanStatus,
		ResourceType:    scanResults.ResourceType,
		ResourceName:    scanResults.ResourceName,
		Compliant:       scanResults.Summary.Compliant,
		FixableCount:    scanResults.Summary.FixableCount,
		ScannerVersion:  scanResults.ScanMetaData.ScannerVersion,
		PoliciesVersion: scanResults.ScanMetaData.PoliciesVersion,
	}

	var vulnDetails []models.VulnerabilityDetails
	var riskFactors []models.RiskFactor

	for _, vuln := range scanResults.Vulnerabilities {
		vulnDetails = append(vulnDetails, models.VulnerabilityDetails{
			VulnID:         vuln.ID,
			ScanID:         scanResults.ScanID,
			Severity:       vuln.Severity,
			CVSS:           vuln.CVSS,
			Status:         vuln.Status,
			PackageName:    vuln.PackageName,
			CurrentVersion: vuln.CurrentVersion,
			FixedVersion:   vuln.FixedVersion,
			Description:    vuln.Description,
			PublishedDate:  vuln.PublishedDate,
			Link:           vuln.Link,
		})

		for _, risk := range vuln.RiskFactors {
			riskFactors = append(riskFactors, models.RiskFactor{
				VulnID:     vuln.ID,
				RiskFactor: risk,
			})
		}
	}

	var severityCounts []models.SeverityCount
	for severity, count := range scanResults.Summary.SeverityCounts {
		severityCounts = append(severityCounts, models.SeverityCount{
			ScanID:   scanResults.ScanID,
			Severity: severity,
			Count:    count,
		})
	}

	var scanningRules []models.ScanningRule
	for _, rule := range scanResults.ScanMetaData.ScanningRules {
		scanningRules = append(scanningRules, models.ScanningRule{
			ScanID: scanResults.ScanID,
			Rule:   rule,
		})
	}

	var excludedPaths []models.ExcludedPath
	for _, path := range scanResults.ScanMetaData.ExcludedPaths {
		excludedPaths = append(excludedPaths, models.ExcludedPath{
			ScanID: scanResults.ScanID,
			Path:   path,
		})
	}
	return scan, vulnDetails, riskFactors, severityCounts, scanningRules, excludedPaths
}
