package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func TransactionHandler(c *gin.Context) {
	db, err := gorm.Open("sqlite", "/tmp/stock.db")
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	tx := db.Begin()
	c.Set("transaction", tx)
	c.Next()

	if err := recover(); err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(500, gin.H{"message": "Could not write to database"})
	}
}
