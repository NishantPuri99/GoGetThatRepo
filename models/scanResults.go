package models

type ScanWrapper struct {
	ScanResult ScanResults `json:"scanResults"`
}

type ScanResults struct {
	ScanID          string          `json:"scan_id"`
	Timestamp       string          `json:"timestamp"`
	ScanStatus      string          `json:"scan_status"`
	ResourceType    string          `json:"resource_type"`
	ResourceName    string          `json:"resource_name"`
	Vulnerabilities []Vulnerability `json:"vulnerabilities"`
	Summary         ScanSummary     `json:"summary"`
	ScanMetaData    ScanMetaData    `json:"scan_metadata"`
}

type Vulnerability struct {
	ID             string   `json:"id"`
	Severity       string   `json:"severity"`
	CVSS           float32  `json:"cvss"`
	Status         string   `json:"status"`
	PackageName    string   `json:"package_name"`
	CurrentVersion string   `json:"current_version"`
	FixedVersion   string   `json:"fixed_version"`
	Description    string   `json:"description"`
	PublishedDate  string   `json:"published_date"`
	Link           string   `json:"link"`
	RiskFactors    []string `json:"risk_factors"`
}

type ScanSummary struct {
	TotalVulnerabilities int            `json:"total_vulnerabilities"`
	SeverityCounts       map[string]int `json:"severity_counts"`
	FixableCount         int            `json:"fixable_count"`
	Compliant            bool           `json:"compliant"`
}

type ScanMetaData struct {
	ScannerVersion  string   `json:"scanner_version"`
	PoliciesVersion string   `json:"policies_version"`
	ScanningRules   []string `json:"scanning_rules"`
	ExcludedPaths   []string `json:"excluded_paths"`
}
