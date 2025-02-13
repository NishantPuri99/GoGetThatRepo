package utils

import (
	"encoding/json"
	"fmt"
	"github-scanner/models"
)

const UNMARSHAL_ERROR_MSG = "error unmarshalling JSON: %v"

func ParseScanResults(content string) ([]models.ScanWrapper, error) {
	var scanWrapper []models.ScanWrapper
	err := json.Unmarshal([]byte(content), &scanWrapper)
	if err != nil {
		return nil, fmt.Errorf(UNMARSHAL_ERROR_MSG, err)
	}
	return scanWrapper, nil
}
