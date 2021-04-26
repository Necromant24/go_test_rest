package api

import (
	"test/DB"
	"test/models"

	"github.com/gin-gonic/gin"
)

func CreateTable(ctx *gin.Context) {
	var table models.CardTable

	ctx.ShouldBindJSON(&table)

	DB.CreateTable(table)
}

func CreateCardList(ctx *gin.Context) {
	var cardList models.CardListDTO

	ctx.ShouldBindJSON(&cardList)

	DB.CreateCardList(cardList)
}

func CreateCard(ctx *gin.Context) {
	var card models.CardDTO

	ctx.ShouldBindJSON(&card)

	DB.CreateCard(card)
}

func GetTableCardLists(ctx *gin.Context) {
	var table models.CardTable

	ctx.ShouldBindQuery(&table)

	ctx.JSON(200, DB.GetTableCardListsById(table.Id))
}
