package DB

import (
	"fmt"

	_ "github.com/lib/pq"
)

var dbInit = `
CREATE TABLE IF NOT EXISTS Card_Tables (
	id SERIAL,
	name varchar(45) NOT NULL,
	PRIMARY KEY (id)
  );

  CREATE TABLE IF NOT EXISTS Card_Lists (
	id SERIAL,
	name varchar(45) NOT NULL,
	PRIMARY KEY (id)
  );


  CREATE TABLE IF NOT EXISTS Card (
	id SERIAL,
	name varchar(45) NOT NULL,
	description TEXT,
	PRIMARY KEY (id)
  );


  CREATE TABLE IF NOT EXISTS cards_to_card_list (
	card_id integer NOT NULL,
	cardlist_id integer NOT NULL
  );

  CREATE TABLE IF NOT EXISTS card_lists_to_card_table (
	cardlist_id integer NOT NULL,
	cardtable_id integer NOT NULL
  );

`

func SeedDb() {

	if db == nil {
		InitDBConnection()
	}

	res := db.MustExec(dbInit)

	fmt.Println(res.RowsAffected())

	tx := db.MustBegin()

	tx.MustExec("INSERT INTO Card_Tables ( name) VALUES ( $1)", "project1")
	tx.MustExec("INSERT INTO Card_Tables ( name) VALUES ($1)", "project2")

	tx.MustExec("INSERT INTO Card_Lists ( name) VALUES ($1)", "Todo")
	tx.MustExec("INSERT INTO Card_Lists ( name) VALUES ($1)", "Todo")
	tx.MustExec("INSERT INTO Card_Lists ( name) VALUES ($1)", "Done")
	tx.MustExec("INSERT INTO Card_Lists ( name) VALUES ($1)", "Done")
	tx.MustExec("INSERT INTO Card_Lists ( name) VALUES ($1)", "Waiting")

	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 2", "Some description1")
	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 3", "Some description2")
	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 4", "Some description3")
	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 5", "Some description4")
	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 6", "Some description5")
	tx.MustExec("INSERT INTO Card ( name, description) VALUES ($1, $2)", "DO 7", "Some description6")

	tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 0, 0)
	tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 0, 2)
	tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 1)
	tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 3)
	tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 4)

	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 0, 0)
	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 0, 1)
	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 1, 2)
	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 2, 3)
	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 3, 4)
	tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 4, 5)

	tx.Commit()

	//tables := GetAllTables()

	//table := GetTableCardListsById(tables[0].Id)

	//fmt.Println(table)

}
