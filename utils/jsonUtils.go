package utils

import (
	"encoding/json"
	"fmt"
	"github-scanner/models"
)

func ParseScanResults(content string) ([]models.ScanWrapper, error) {
	var scanWrapper []models.ScanWrapper
	err := json.Unmarshal([]byte(content), &scanWrapper)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}
	return scanWrapper, nil
}
