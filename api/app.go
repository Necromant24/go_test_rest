package api

import (
	"test/DB"

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

		apiArea.GET("/tableCardLists", GetTableCardLists)

		apiArea.POST("/table", CreateTable)

		apiArea.POST("/cardlist", CreateCardList)

		apiArea.POST("/card", CreateCard)

	}

	r.Run()
}
