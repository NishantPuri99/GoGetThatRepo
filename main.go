package main

import (
	"github-scanner/api"
	"github-scanner/db"
	"github-scanner/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(string(utils.DB_INIT_FAILED_MESSAGE), err)
	} else {
		r := gin.Default()
		r.POST("/scan", func(c *gin.Context) {
			api.ScanRepo(c, database)
		})
		r.POST("/query", func(c *gin.Context) {
			api.QueryWithFilter(c, database)
		})
		r.Run(":8080")
	}
}
