package models

type CardList struct {
	Id    int    `db:"id"`
	Name  string `json:"name" db:"name"`
	Cards []Card `json:"cards"`
}

type CardListDTO struct {
	Name    string `json:"name" db:"name" binding:"required"`
	TableId int    `json:"table_id"`
}
