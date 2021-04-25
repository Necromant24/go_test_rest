package models

type CardTable struct {
	Id        int        `db:"id" form:"id"`
	Name      string     `json:"name" db:"name" form:"name"`
	CardLists []CardList `json:"cardLists"`
}
