package DB

import (
	"strconv"

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
