package api

import (
	"fmt"

	"github-scanner/db"
	"github-scanner/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RequestData struct {
	Repo  string   `json:"repo"`
	Files []string `json:"files"`
}

func internalServerErrorHandler(context *gin.Context, errors []error, fetchErrors bool) {
	var errOutput []string
	for _, err := range errors {
		errString := fmt.Sprintf("%v", err)
		errOutput = append(errOutput, string(errString))
	}
	var errorMessage string
	if fetchErrors {
		errorMessage = string(utils.FILE_FETCH_ERROR_MESSAGE)
	} else {
		errorMessage = string(utils.DB_ADD_ERROR_MESSAGE)
	}
	context.JSON(http.StatusInternalServerError, gin.H{
		string(utils.ERROR):   errorMessage,
		string(utils.DETAILS): errOutput,
	})
}

func ScanRepo(context *gin.Context, database *gorm.DB) {
	var requestData RequestData

	if err := context.ShouldBindJSON(&requestData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"Invalid Request": "Please provide a valid JSON payload in the format: { ‘repo’: ‘’, ‘files’: [’’, ‘’, …] }."})
		return
	}

	var urls []string

	for _, filename := range requestData.Files {
		fileURL := utils.GetRawGitHubURL(requestData.Repo, filename)
		fmt.Printf("[SCAN API] Fetching file: %s\n", fileURL)
		urls = append(urls, fileURL)
	}

	fileContents, errors := utils.FetchMultipleFiles(urls)
	if errors != nil {
		internalServerErrorHandler(context, errors, true)
		return
	}

	var dbErrors []error
	fileCount := 0
	for url, content := range fileContents {
		fmt.Printf("[SCAN API] Processed file from %s\n", url)
		scanWrapper, err := utils.ParseScanResults(content)
		if err != nil {
			log.Printf("[SCAN API] Error Parsing file %s: %v", url, err)
			dbErrors = append(dbErrors, err)
			continue
		} else {
			for _, result := range scanWrapper {
				err := db.SaveScan(database, result)
				if err != nil {
					log.Printf("[SCAN API] Failed to store scan results: %v", err)
					dbErrors = append(dbErrors, err)
				} else {
					fileCount++
				}
			}
		}
	}
	if dbErrors != nil {
		internalServerErrorHandler(context, dbErrors, false)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		string(utils.FILES_ADDED_SUCCESS_MESSAGE): fileCount,
	})
}
