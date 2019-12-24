package api

import (
	"github.com/JHeinzde/stock-data-go/backend/database"
	"github.com/gin-gonic/gin"
	"log"
)

func GetActiveStocks(c *gin.Context) {
	c.JSON(200, database.GetAllStocks(c))
}

func AddNewImportTarget(c *gin.Context) {
	var body = StockSymbol{}
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(500, gin.H{"message": "Could not parse json body"})
		return
	}

	symbol := body.Symbol
	database.InsertStock(symbol, c)

	c.JSON(200, gin.H{"message": "Success"})
}

func StartAndServeAPI() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(TransactionHandler)

	r.GET("/stocks", GetActiveStocks)
	r.POST("/stocks", AddNewImportTarget)
	log.Fatalln(r.Run("0.0.0.0:8089"))
}
