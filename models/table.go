package models

type CardTable struct {
	Id               int        `db:"id" form:"id"`
	Name             string     `json:"name" db:"name" form:"name" binding:"required"`
	BriefDescription string     `json:"briefDescription" db:"brief_description"`
	CardLists        []CardList `json:"cardLists"`
}

type CTableFilter struct {
	Name string `json:"name"`
}

type DbFullCardTable struct {
	CardListId   int    `db:"card_list_id"`
	CardListName string `db:"card_list_name"`

	CardId          int    `db:"card_id"`
	CardName        string `db:"card_name"`
	CardDescription string `db:"card_description"`
}
