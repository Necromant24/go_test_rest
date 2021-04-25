package models

type Card struct {
	Id          int    `db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}
