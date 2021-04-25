package main

import (
	"test/some"

	"fmt"
	"test/DB"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123"
	dbname   = "postgres"
)

func main() {
	fmt.Println("go brrrrrrrrrr.")
	fmt.Println(some.Say("hi"))

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	//conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	fmt.Println(psqlInfo)

	DB.SeedDb()

}
