package models

type CardTable struct {
	Id        int        `db:"id"`
	Name      string     `json:"name" db:"name"`
	CardLists []CardList `json:"cardLists"`
}
