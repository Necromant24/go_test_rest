package actions

import (
	"strconv"
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

func GetAllTables(ctx *gin.Context) {
	tables := CardTables.GetAllTables()

	ctx.JSON(200, tables)
}

func GetFullTable(ctx *gin.Context) {
	tableId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}

	table, _ := CardTables.GetFullTable(int(tableId))

	ctx.JSON(200, table)
}

func UpdateCard(ctx *gin.Context) {
	var card models.Card

	err := ctx.ShouldBind(&card)
	if err != nil {
		panic(err)
	}

	err = Cards.UpdateCard(card)

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func UpdateCardList(ctx *gin.Context) {
	var cardList models.CardList

	err := ctx.ShouldBind(&cardList)
	if err != nil {
		panic(err)
	}

	err = CardLists.UpdateCList(cardList)

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func UpdateTable(ctx *gin.Context) {
	var table models.CardTable

	err := ctx.ShouldBind(&table)
	if err != nil {
		panic(err)
	}

	err = CardTables.UpdateTable(table)

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteTable(ctx *gin.Context) {
	tableId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = CardTables.DeleteTable(int(tableId))

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteCardList(ctx *gin.Context) {
	cListId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = CardLists.DeleteCList(int(cListId))

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteCard(ctx *gin.Context) {
	card, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = Cards.DeleteCard(int(card))

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func CreateCardLink(ctx *gin.Context) {
	var cardLink models.CardLink

	err := ctx.ShouldBind(&cardLink)

	err = Cards.CreateLink(cardLink)

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}

func DeleteCardLink(ctx *gin.Context) {
	keyId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = Cards.DeleteLink(int(keyId))

	if err != nil {
		panic(err)
	}

	ctx.JSON(200, gin.H{
		"message": "ok",
	})
}
