package Cards

import (
	"test/DB"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDBConnection() {
	db = DB.GetDBConnection()
}

func CreateCard(card models.CardDTO) error {
	_, err := db.Exec("INSERT INTO Card (name, description, card_list_id) VALUES ($1, $2, $3);",
		card.Name, card.Description, card.CardListId)

	return err
}

func UpdateCard(card models.Card) error {
	_, err := db.Exec("UPDATE Card set name = $1, description = $2 WHERE id = $3;",
		card.Name, card.Description, card.Id)

	return err
}

func DeleteCard(cardId int) error {
	_, err := db.Exec("DELETE FROM Card WHERE id = $1", cardId)

	return err
}

func CreateLink(link models.CardLink) error {
	_, err := db.Exec("INSERT INTO card_to_card (key_id, value_id) VALUES($1, $2);",
		link.KeyId, link.ValueId)

	return err
}

func DeleteLink(keyId int) error {
	_, err := db.Exec("DELETE FROM card_to_card WHERE id = $1", keyId)

	return err
}
