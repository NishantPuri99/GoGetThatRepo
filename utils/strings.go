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
const INVALID_REQUEST = "Invalid Request"
const BAD_REQUEST_ERROR_MESSAGE = "Please provide a valid JSON payload in the format: { 'repo': '', 'files': ['', '', …] }."

// ../api/query.go
const INVALID_JSON_BAD_REQUEST_ERROR_MESSAGE = "Please provide a valid JSON payload in the format: { 'filters': {'severity':'CRITICAL'} }."
const INVALID_SEV_VALUE_ERROR_MESSAGE = "Invalid severity value. Allowed: CRITICAL, HIGH, MEDIUM, LOW"
const DB_QUERY_ERROR = "Database Query Error: %v\n"
const FAILED_TO_FETCH_VULNS = "Failed to fetch vulnerabilities"

// [SCAN API] messages
const FETCH_FILE_MESSAGE = "[SCAN API] Fetching file: %s\n"
const PROCESSED_FILE_MESSAGE = "[SCAN API] Processed file from %s\n"
const PARSING_ERROR_MESSAGE = "[SCAN API] Error Parsing file %s: %v"
const STORE_SCAN_ERROR_MESSAGE = "[SCAN API] Failed to store scan results: %v"

// ../db/scanDb.go
const DB_SCAN_INSERT_ERROR = "[DB] Error inserting scan: %v"
const DB_VULN_INSERT_ERROR = "[DB] Error inserting vulnerabilities: %v"
const DB_RISK_FACTOR_INSERT_ERROR = "[DB] Error inserting risk factors: %v"
const DB_SEV_COUNT_INSERT_ERROR = "[DB] Error inserting severity counts: %v"
const DB_SCAN_RULES_INSERT_ERROR = "[DB] Error inserting scanning rules: %v"
const DB_EXC_PATH_INSERT_ERROR = "[DB] Error inserting excluded paths: %v"
const DB_SUCCESS_SAVE = "[DB] Scan successfully saved!"

// ../db/db.go
const DB_CONNECT_FAIL = "Failed to connect to database:"
const DB_MIGRATE_FAIL = "Failed to migrate database:"
const DB_TABLE_CREATION_SUCCESS = "Database and tables created successfully."
