package api

import (
	"test/DB"
	"test/models"

	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiArea := r.Group("/api")
	{
		apiArea.GET("/some", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "some",
			})
		})

		apiArea.GET("/tables", func(c *gin.Context) {
			c.JSON(200, DB.GetAllTables())
		})

		apiArea.GET("/fulltable", func(c *gin.Context) {

			var table models.CardTable
			c.ShouldBindQuery(&table)

			c.JSON(200, DB.GetFullTableByTable(table))
		})

		apiArea.POST("/table", CreateTable)

		apiArea.POST("/cardlist", CreateCardList)

	}

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
