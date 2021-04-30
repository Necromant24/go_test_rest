package actions

import (
	"test/DB/CardLists"
	"test/DB/CardTables"
	"test/DB/Cards"
	"test/models"

	"github.com/gin-gonic/gin"
)

// TODO: add all actions

func CreateTable(ctx *gin.Context) {
	var table models.CardTable

	ctx.ShouldBindJSON(&table)

	CardTables.CreateTable(table)
}

func CreateCardList(ctx *gin.Context) {
	var cardList models.CardListDTO

	ctx.ShouldBindJSON(&cardList)

	CardLists.CreateCardList(cardList)
}

func CreateCard(ctx *gin.Context) {
	var card models.CardDTO

	ctx.ShouldBindJSON(&card)

	Cards.CreateCard(card)
}
