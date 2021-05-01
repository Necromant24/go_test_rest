package models

type Card struct {
	Id          int    `db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
}

type CardDTO struct {
	Name        string `json:"name" db:"name" binding:"required"`
	Description string `json:"description" db:"description"`
	CardListId  int    `json:"cardListId"`
}

type CardLink struct {
	KeyId   int `db:"key_id"`
	ValueId int `db:"value_id"`
}
