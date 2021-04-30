package CardLists

import (
	"test/DB"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDBConnection() {
	db = DB.GetDBConnection()
}

func CreateCardList(cList models.CardListDTO) error {
	_, err := db.Exec("INSERT INTO Card_Lists (name, table_id) VALUES ($1, $2) ;", cList.Name, cList.TableId)

	return err
}

func UpdateCList(cList models.CardList) error {
	_, err := db.Exec("UPDATE card_lists SET name = '$1' WHERE id = $2", cList.Name, cList.Id)

	return err
}

func DeleteCList(cListId int) error {
	_, err := db.Exec("DELETE FROM card_lists WHERE id = $1", cListId)

	return err
}
