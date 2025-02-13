package utils

// ../main.go
const DB_INIT_FAILED_MESSAGE = "Failed to initialize DB:"

// ../api/scan.go
// Response Messages
const FILES_ADDED_SUCCESS_MESSAGE = "Files added successfully"
const FILE_FETCH_ERROR_MESSAGE = "Some files could not be fetched."
const DB_ADD_ERROR_MESSAGE = "Entries could not be added to Database"
const ERROR = "error"
const DETAILS = "details"

// ../db/scanDb.go
const DB_SCAN_INSERT_ERROR = "[DB] Error inserting scan: %v"
const DB_VULN_INSERT_ERROR = "[DB] Error inserting vulnerabilities: %v"
const DB_RISK_FACTOR_INSERT_ERROR = "[DB] Error inserting risk factors: %v"
const DB_SEV_COUNT_INSERT_ERROR = "[DB] Error inserting severity counts: %v"
const DB_SCAN_RULES_INSERT_ERROR = "[DB] Error inserting scanning rules: %v"
const DB_EXC_PATH_INSERT_ERROR = "[DB] Error inserting excluded paths: %v"
const DB_SUCCESS_SAVE = "[DB] Scan successfully saved!"
