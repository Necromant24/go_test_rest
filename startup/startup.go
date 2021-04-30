package startup

import (
	"test/DB/CardLists"
	"test/DB/CardTables"
	"test/DB/Cards"
	"test/api"
)

func LoadApp() {

	initAllDatabases()
	api.Run()
}

func initAllDatabases() {

	CardLists.InitDBConnection()
	Cards.InitDBConnection()
	CardTables.InitDBConnection()
}
