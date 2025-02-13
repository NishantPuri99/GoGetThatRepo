package models

type Scan struct {
	ScanID          string `gorm:"primaryKey"`
	Timestamp       string
	ScanStatus      string
	ResourceType    string
	ResourceName    string
	ScannerVersion  string
	PoliciesVersion string
	Compliant       bool
	FixableCount    int
}

type VulnerabilityDetails struct {
	VulnID         string `gorm:"primaryKey"`
	ScanID         string `gorm:"index"`
	Severity       string
	CVSS           float32
	Status         string
	PackageName    string
	CurrentVersion string
	FixedVersion   string
	Description    string
	PublishedDate  string
	Link           string
}

type RiskFactor struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	VulnID     string `gorm:"index"`
	RiskFactor string
}

type SeverityCount struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	ScanID   string `gorm:"index"`
	Severity string
	Count    int
}

type ScanningRule struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	ScanID string `gorm:"index"`
	Rule   string
}

type ExcludedPath struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	ScanID string `gorm:"index"`
	Path   string
}
