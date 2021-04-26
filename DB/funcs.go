package DB

import (
	"strconv"
	"test/models"

	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

var db *sqlx.DB

func CreateTable(table models.CardTable) error {
	_, err := db.Exec("INSERT INTO Card_Tables (name) VALUES ($1)", table.Name)

	return err
}

func CreateCardList(cList models.CardListDTO) error {
	//res, err := db.Exec("INSERT INTO Card_Lists (name) VALUES ($1) RETURNING id;", cList.Name)

	var lastId int

	err := db.Get(&lastId, "INSERT INTO Card_Lists (name) VALUES ($1) RETURNING id;", cList.Name)

	_, err = db.Exec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id ) VALUES ($1, $2);", cList.TableId, lastId)

	return err
}

func CreateCard(card models.CardDTO) error {
	var lastId int

	err := db.Get(&lastId, "INSERT INTO Card (name, description) VALUES ($1, $2) RETURNING id;", card.Name, card.Description)

	_, err = db.Exec("INSERT INTO cards_to_card_list (card_id, cardlist_id ) VALUES ($1, $2);", lastId, card.CardListId)

	return err
}

func InitDBConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbConn, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	db = dbConn
}

func convertIntArrToInt64(items []int) []int64 {
	var converted = make([]int64, len(items))

	for i, el := range items {
		converted[i] = int64(el)
	}

	return converted
}

func buildIdsString(Ids []int) string {
	var ids = convertIntArrToInt64(Ids)

	length := len(ids)

	if length == 0 {
		return ""
	} else if length == 1 {
		return strconv.FormatInt(ids[0], 10)
	} else {

		builder := strconv.FormatInt(ids[0], 10)

		for i := 1; i < length; i++ {
			builder += "," + strconv.FormatInt(ids[i], 10)
		}

		return builder
	}

}

func GetAllTables() []models.CardTable {

	tables := []models.CardTable{}
	err := db.Select(&tables, "SELECT id, name FROM card_tables ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	return tables
}

//GetCListsToCTableLinks
func GetCListToCTableIds(tableId int) ([]int, []models.CListsToCTableLink) {

	cListsToCTableLinks := []models.CListsToCTableLink{}
	err := db.Select(&cListsToCTableLinks, "SELECT cardlist_id, cardtable_id FROM card_lists_to_card_table WHERE cardtable_id=$1", tableId)
	if err != nil {
		panic(err)
	}

	cListsToCTableLinksCount := len(cListsToCTableLinks)
	cListIds := make([]int, cListsToCTableLinksCount)

	for i, el := range cListsToCTableLinks {
		cListIds[i] = el.CListId
	}

	return cListIds, cListsToCTableLinks
}

func GetCListsByIds(cListIds []int) []models.CardList {

	cardLists := []models.CardList{}
	err := db.Select(&cardLists, "SELECT id, name FROM Card_Lists WHERE id IN("+buildIdsString(cListIds)+");")
	if err != nil {
		panic(err)
	}

	return cardLists
}

func GetCardsToClistLinksByIds(cardListsIds []int) []models.CardsToClistLink {
	cardsToClistLinks := []models.CardsToClistLink{}
	err := db.Select(&cardsToClistLinks, "SELECT cardlist_id, card_id FROM cards_to_card_list WHERE cardlist_id IN("+buildIdsString(cardListsIds)+");")
	if err != nil {
		panic(err)
	}

	return cardsToClistLinks
}

func GetAssignedCardsToCLists(cardLists []models.CardList) []models.CardList {
	cardListsIds := make([]int, len(cardLists))

	for i, el := range cardLists {
		cardListsIds[i] = el.Id
	}

	cardsToClistLinks := GetCardsToClistLinksByIds(cardListsIds)

	if len(cardsToClistLinks) == 0 {
		return cardLists
	}

	var cardsIds []int

	cardsIds = make([]int, len(cardsToClistLinks))

	for i, el := range cardsToClistLinks {
		cardsIds[i] = el.CardId
	}

	var cards []models.Card
	err := db.Select(&cards, "SELECT id,name,description FROM Card WHERE id IN("+buildIdsString(cardsIds)+")")
	if err != nil {
		panic(err)
	}

	// assign cards to cardLists
	for i, el := range cardLists {

		for _, link := range cardsToClistLinks {
			if link.CListId == el.Id {

				var card models.Card
				for _, item := range cards {
					if item.Id == link.CardId {
						card = item
						break
					}
				}

				cardLists[i].Cards = append(cardLists[i].Cards, card)
			}
		}
	}

	return cardLists
}

// only assign CardLists with assigned Cards to CardLists
func GetTableCardListsById(tableId int) []models.CardList {

	cListIds, _ := GetCListToCTableIds(tableId)

	cardLists := GetCListsByIds(cListIds)

	cardLists = GetAssignedCardsToCLists(cardLists)

	return cardLists
}
