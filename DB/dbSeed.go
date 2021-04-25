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

	table1_Id := tables[0].Id

	cListIds, _ := GetCListToCTableIds(db, table1_Id)

	// cListsToCTableLinks := []models.CListsToCTableLink{}
	// db.Select(&cListsToCTableLinks, "SELECT cardlist_id, cardtable_id FROM card_lists_to_card_table WHERE cardtable_id=$1", table1_Id)

	// cListsToCTableLinksCount := len(cListsToCTableLinks)
	// cListIds := make([]int, cListsToCTableLinksCount)

	// for i, el := range cListsToCTableLinks {
	// 	cListIds[i] = el.CListId
	// }

	//cListIdsString := buildIdsString(cListIds)

	// cardLists := []models.CardList{}
	// //query := "SELECT id, name FROM Card_Lists WHERE id IN(" + buildIdsString(cListIds) + ");"
	// //err = db.Select(&cardLists, "SELECT id, name FROM Card_Lists WHERE id IN($1)", cListIds)
	// err = db.Select(&cardLists, "SELECT id, name FROM Card_Lists WHERE id IN("+buildIdsString(cListIds)+");")

	cardLists := GetCListsByIds(db, cListIds)

	//cardListCount := len(cardLists)
	// cardListsIds := make([]int, len(cardLists))

	// for i, el := range cardLists {
	// 	cardListsIds[i] = el.Id
	// }

	// //cardListsIdsString := buildIdsString(cardListsIds)

	// // cardsToClistLinks := []models.CardsToClistLink{}
	// // //query = "SELECT cardlist_id, card_id FROM cards_to_card_list WHERE cardlist_id IN(" + buildIdsString(cardListsIds) + ");"
	// // //err = db.Select(&cardsToClistLinks, "SELECT cardlist_id, card_id FROM cards_to_card_list WHERE cardlist_id IN($1)", cardListsIdsString)
	// // err = db.Select(&cardsToClistLinks, "SELECT cardlist_id, card_id FROM cards_to_card_list WHERE cardlist_id IN("+buildIdsString(cardListsIds)+");")

	// cardsToClistLinks := GetCardsToClistLinksByIds(db, cardListsIds)
	// //cardsToClistLinksCount := len(cardsToClistLinks)

	// var cardsIds []int

	// cardsIds = make([]int, len(cardsToClistLinks))

	// for i, el := range cardsToClistLinks {
	// 	cardsIds[i] = el.CardId
	// }

	// //cardsIdsString := buildIdsString(cardsIds)

	// var cards []models.Card
	// //query = "SELECT id,name,description FROM Card WHERE id IN(" + buildIdsString(cardsIds) + ")"
	// //err = db.Select(&cards, "SELECT id,name,description FROM Cards WHERE id IN($1)", cardsIdsString)
	// err = db.Select(&cards, "SELECT id,name,description FROM Card WHERE id IN("+buildIdsString(cardsIds)+")")

	// // assign cards to cardLists
	// for i, el := range cardLists {

	// 	for _, link := range cardsToClistLinks {
	// 		if link.CListId == el.Id {

	// 			var card models.Card
	// 			for _, item := range cards {
	// 				if item.Id == link.CardId {
	// 					card = item
	// 					break
	// 				}
	// 			}

	// 			cardLists[i].Cards = append(cardLists[i].Cards, card)
	// 		}
	// 	}
	// }

	cardLists = GetAssignedCardsToCLists(db, cardLists)

	// assign cardLists to tables
	table := tables[0]

	table.CardLists = cardLists

	// for _, link := range cListsToCTableLinks {
	// 	if link.CTableId == table.Id {

	// 		// get CardList item assigned to current table
	// 		var cListItem models.CardList
	// 		for _, item := range cardLists {
	// 			if item.Id == link.CListId {
	// 				cListItem = item
	// 				break
	// 			}
	// 		}

	// 		table.CardLists = append(table.CardLists, cListItem)
	// 	}
	// }

	fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA!!!!!!!!!!")

	//!!!!!!!
	//!!!!!!!
	//!!!!!!!

	//!!!!!!!
	//!!!!!!!

	//!!!!!!!
	//!!!!!!!
	//!!!!!!!//!!!!!!!
	//!!!!!!!
	//!!!!!!!

	//!!!!!!!//!!!!!!!

	// Query the database, storing results in a []Person (wrapped in []interface{})
	people := []Person{}
	db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}
	// Person{FirstName:"John", LastName:"Doe", Email:"johndoeDNE@gmail.net"}

	// You can also get a single result, a la QueryRow
	jason = Person{}
	err = db.Get(&jason, "SELECT * FROM person WHERE first_name=$1", "Jason")
	fmt.Printf("%#v\n", jason)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []Place{}
	err = db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}

	// Loop through rows using only one struct
	place := Place{}
	rows, err := db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, err = db.NamedExec(`INSERT INTO person (first_name,last_name,email) VALUES (:first,:last,:email)`,
		map[string]interface{}{
			"first": "Bin",
			"last":  "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, err = db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)

	// batch insert

	// batch insert with structs
	personStructs := []Person{
		{FirstName: "Ardie", LastName: "Savea", Email: "asavea@ab.co.nz"},
		{FirstName: "Sonny Bill", LastName: "Williams", Email: "sbw@ab.co.nz"},
		{FirstName: "Ngani", LastName: "Laumape", Email: "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personStructs)

	// batch insert with maps
	personMaps := []map[string]interface{}{
		{"first_name": "Ardie", "last_name": "Savea", "email": "asavea@ab.co.nz"},
		{"first_name": "Sonny Bill", "last_name": "Williams", "email": "sbw@ab.co.nz"},
		{"first_name": "Ngani", "last_name": "Laumape", "email": "nlaumape@ab.co.nz"},
	}

	_, err = db.NamedExec(`INSERT INTO person (first_name, last_name, email)
        VALUES (:first_name, :last_name, :email)`, personMaps)
}
