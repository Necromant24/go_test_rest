package models

type CardList struct {
	Id    int    `db:"id"`
	Name  string `json:"name" db:"name"`
	Cards []Card `json:"cards"`
}
