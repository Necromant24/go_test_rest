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

		apiArea.GET("/table", actions.GetAllTables)
		apiArea.GET("/table/:id", actions.GetFullTable)

		apiArea.POST("/table", actions.CreateTable)
		apiArea.POST("/cardlist", actions.CreateCardList)
		apiArea.POST("/card", actions.CreateCard)
		apiArea.POST("cardLink", actions.CreateCardLink)

		apiArea.DELETE("/table/:id", actions.DeleteTable)
		apiArea.DELETE("/cardList/:id", actions.DeleteCardList)
		apiArea.DELETE("/card/:id", actions.DeleteCard)
		apiArea.DELETE("/cardLink", actions.DeleteCardLink)

		apiArea.PATCH("/table/:id", actions.UpdateTable)
		apiArea.PATCH("/cardList/:id", actions.UpdateCardList)
		apiArea.PATCH("/card/:id", actions.UpdateCard)

	}

	r.Run()
}
