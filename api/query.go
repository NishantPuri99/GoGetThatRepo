package api

import (
	"github-scanner/db"
	"github-scanner/utils"
	"log"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QueryStruct struct {
	Filters struct {
		Severity string `json:"severity"`
	} `json:"filters"`
}

func QueryWithFilter(context *gin.Context, database *gorm.DB) {
	var filterQuery QueryStruct

	if err := context.ShouldBindJSON(&filterQuery); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{string(utils.INVALID_REQUEST): utils.INVALID_JSON_BAD_REQUEST_ERROR_MESSAGE})
		return
	}

	if reflect.ValueOf(filterQuery).IsZero() {
		context.JSON(http.StatusBadRequest, gin.H{string(utils.INVALID_REQUEST): utils.INVALID_JSON_BAD_REQUEST_ERROR_MESSAGE})
		return
	}

	validSeverities := map[string]bool{"CRITICAL": true, "HIGH": true, "MEDIUM": true, "LOW": true}
	if !validSeverities[filterQuery.Filters.Severity] {
		context.JSON(http.StatusBadRequest, gin.H{string(utils.INVALID_REQUEST): utils.INVALID_SEV_VALUE_ERROR_MESSAGE})
		return
	}
	returnStr, err := db.QueryVulns(database, filterQuery.Filters.Severity)
	if err != nil {
		log.Printf(utils.DB_QUERY_ERROR, err)
		context.JSON(http.StatusInternalServerError, gin.H{string(utils.ERROR): utils.FAILED_TO_FETCH_VULNS})
		return
	}
	context.JSON(http.StatusOK, returnStr)
}
