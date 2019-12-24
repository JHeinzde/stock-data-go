package api

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetActiveStocks(c *gin.Context) {
	c.JSON(200, LoadFile())
}

func AddNewImportTarget(c *gin.Context) {
	var body = StockSymbol{}
	err := c.BindJSON(&body)

	if err != nil {
		c.JSON(500, gin.H{"message": "could not parse json body"})
		return
	}

	symbol := body.Symbol
	AddLineToFile(symbol)
	log.Println("WTF")

	c.JSON(200, gin.H{"message": "success"})
}

func StartAndServeAPI() {
	r := gin.Default()

	r.GET("/stocks", GetActiveStocks)
	r.POST("/stocks", AddNewImportTarget)
	log.Fatalln(r.Run("0.0.0.0:8089"))
}
