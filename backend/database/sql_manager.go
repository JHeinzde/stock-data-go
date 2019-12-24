package database

import "github.com/gin-gonic/gin"
import "github.com/jinzhu/gorm"

func InsertUser(user User, c *gin.Context) {
	tx := getTransactionFromContext(c)

	tx.Create(&user)
}

func InsertStock(symbol string, c *gin.Context) {
	stock := Stock{Symbol: symbol}
	tx := getTransactionFromContext(c)

	tx.Create(&stock)
}

func RemoveUser(user User, c *gin.Context) {
	tx := getTransactionFromContext(c)

	tx.Delete(&user)
}

func RemoveStock(symbol string, c *gin.Context) {
	stock := Stock{Symbol: symbol}
	tx := getTransactionFromContext(c)

	tx.Delete(&stock)
}

func GetAllStocks(c *gin.Context) *[]Stock {
	tx := getTransactionFromContext(c)
	stocks := make([]Stock, 0)

	tx.Find(&stocks)

	return &stocks
}

func getTransactionFromContext(c *gin.Context) *gorm.DB {
	tx := c.MustGet("transaction")
	return tx.(*gorm.DB)
}
