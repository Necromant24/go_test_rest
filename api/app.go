package api

import (
	"github.com/gin-gonic/gin"

	"test/api/actions"
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
			c.JSON(200, actions.GetAllTables())
		})

		apiArea.GET("/tableCardLists", actions.GetTableCardLists)

		apiArea.POST("/table", actions.CreateTable)

		apiArea.POST("/cardlist", actions.CreateCardList)

		apiArea.POST("/card", actions.CreateCard)

	}

	r.Run()
}
