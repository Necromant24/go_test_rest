package DB

import (
	"log"
	"test/config"

	"github.com/jmoiron/sqlx"
)

func GetDBConnection() *sqlx.DB {

	dbConn, err := sqlx.Connect("postgres", config.GetDbConnString())
	if err != nil {
		log.Fatalln(err)
	}

	return dbConn
}
