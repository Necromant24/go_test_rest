package CardTables

import (
	"test/DB"
	"test/models"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func InitDBConnection() {
	db = DB.GetDBConnection()
}

func GetAllTables() []models.CardTable {

	tables := []models.CardTable{}
	err := db.Select(&tables, "SELECT id, name FROM card_tables ORDER BY id ASC")
	if err != nil {
		panic(err)
	}

	return tables
}

func GetTables(filter models.CTableFilter) []models.CardTable {
	tables := []models.CardTable{}
	err := db.Select(&tables, "SELECT id, name FROM card_tables WHERE name LIKE '%$1%' ", filter.Name)
	if err != nil {
		panic(err)
	}

	return tables
}

func GetTable(tableId int) (models.CardTable, error) {
	var table models.CardTable

	err := db.Get(&table, "SELECT * FROM card_tables WHERE id = $1", tableId)

	return table, err
}

var fullTableSelect = `

SELECT 

	cl.Id AS card_list_id,
	cl.name AS card_list_name,
	
	card.Id AS card_id,
	card.name AS card_name


FROM card_tables AS ct 

LEFT JOIN card_lists AS cl ON ct.Id = cl.table_id  

LEFT JOIN card AS card ON cl.Id = card.card_list_id

WHERE ct.Id = $1

`

func GetFullTable(tableId int) (models.CardTable, error) {

	var dbTable []models.DbFullCardTable

	var table models.CardTable

	err := db.Get(&table, "SELECT * FROM card_tables WHERE id = $1", tableId)

	err = db.Select(&dbTable, fullTableSelect, tableId)

	cListMap := make(map[int][]models.Card)

	// assign cards to card_lists
	for _, item := range dbTable {
		cListMap[item.CardListId] = append(cListMap[item.CardListId],
			models.Card{Id: item.CardId, Name: item.CardName, Description: item.CardDescription})
	}

	var cLists []models.CardList

	// make  card_lists with joined cards
	for _, item := range dbTable {
		cLists = append(cLists,
			models.CardList{Id: item.CardListId, Name: item.CardListName, Cards: cListMap[item.CardListId]})
	}

	table.CardLists = cLists

	return table, err
}

func CreateTable(table models.CardTable) error {
	_, err := db.Exec("INSERT INTO Card_Tables (name, brief_description) VALUES ($1, $2)",
		table.Name, table.BriefDescription)

	return err
}

func UpdateTable(table models.CardTable) error {
	_, err := db.Exec("UPDATE Card_Tables SET name = '$1', brief_description = '$2' WHERE id = $3",
		table.Name, table.BriefDescription, table.Id)

	return err
}

func DeleteTable(tableId int) error {
	_, err := db.Exec("DELETE FROM Card_Tables WHERE id = $1;", tableId)

	return err
}
