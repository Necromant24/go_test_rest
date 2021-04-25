package DB

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

var dbInit = `
CREATE TABLE IF NOT EXISTS Card_Tables (
	id integer NOT NULL,
	name varchar(45) NOT NULL,
	PRIMARY KEY (id)
  );

  CREATE TABLE IF NOT EXISTS Card_Lists (
	id integer NOT NULL,
	name varchar(45) NOT NULL,
	PRIMARY KEY (id)
  );


  CREATE TABLE IF NOT EXISTS Card (
	id integer NOT NULL,
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

var dbInit1 = `
CREATE TABLE IF NOT EXISTS TOST1_to (
	id SERIAL NOT NULL,
	name varchar(45) NOT NULL,
	PRIMARY KEY (id)
  );

`

var dbInit2 = `
CREATE TABLE IF NOT EXISTS cards_to_card_list (
	id SERIAL PRIMARY KEY NOT NULL,
	card_id integer NOT NULL,
	cardlist_id integer NOT NULL,
  );

  CREATE TABLE IF NOT EXISTS card_lists_to_card_table (
	id SERIAL PRIMARY KEY NOT NULL,
	cardlist_id integer NOT NULL,
	cardtable_id integer NOT NULL,
  );

`

func SeedDb() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatalln(err)
	}

	// exec the schema or fail; multi-statement Exec behavior varies between
	// database drivers;  pq will exec them all, sqlite3 won't, ymmv
	res := db.MustExec(dbInit)

	fmt.Println(res.RowsAffected())

	//tx := db.MustBegin()

	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	// tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	// tx.MustExec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	// tx.MustExec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")

	// tx.MustExec("INSERT INTO Card_Tables (id, name) VALUES ($1, $2)", 0, "project1")
	// tx.MustExec("INSERT INTO Card_Tables (id, name) VALUES ($1, $2)", 1, "project2")

	// tx.MustExec("INSERT INTO Card_Lists (id, name) VALUES ($1, $2)", 0, "Todo")
	// tx.MustExec("INSERT INTO Card_Lists (id, name) VALUES ($1, $2)", 1, "Todo")
	// tx.MustExec("INSERT INTO Card_Lists (id, name) VALUES ($1, $2)", 2, "Done")
	// tx.MustExec("INSERT INTO Card_Lists (id, name) VALUES ($1, $2)", 3, "Done")
	// tx.MustExec("INSERT INTO Card_Lists (id, name) VALUES ($1, $2)", 4, "Waiting")

	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 0, "DO 2", "Some description1")
	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 1, "DO 3", "Some description2")
	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 2, "DO 4", "Some description3")
	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 3, "DO 5", "Some description4")
	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 4, "DO 6", "Some description5")
	// tx.MustExec("INSERT INTO Card (id, name, description) VALUES ($1, $2, $3)", 5, "DO 7", "Some description6")

	// tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id ) VALUES ($1, $2)", 0, 0)
	// tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 0, 2)
	// tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 1)
	// tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 3)
	// tx.MustExec("INSERT INTO card_lists_to_card_table (cardtable_id, cardlist_id) VALUES ($1, $2)", 1, 4)

	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 0, 0)
	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 0, 1)
	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 1, 2)
	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 2, 3)
	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 3, 4)
	// tx.MustExec("INSERT INTO cards_to_card_list (cardlist_id, card_id) VALUES ($1, $2)", 4, 5)

	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	//tx.NamedExec("INSERT INTO person (first_name, last_name, email) VALUES (:first_name, :last_name, :email)", &Person{"Jane", "Citizen", "jane.citzen@example.com"})
	//tx.Commit()

	// tables := []models.CardTable{}
	// err = db.Select(&tables, "SELECT id, name FROM card_tables ORDER BY id ASC")

	tables := GetAllTables(db)

	// cListIds, _ := GetCListToCTableIds(db, table1_Id)

	// cardLists := GetCListsByIds(db, cListIds)

	// cardLists = GetAssignedCardsToCLists(db, cardLists)

	// // assign cardLists to tables
	// table := tables[0]

	// table.CardLists = cardLists

	table := GetFullTableByTable(db, tables[0])

	fmt.Println(table)

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA!!!!!!!!!!")

}
