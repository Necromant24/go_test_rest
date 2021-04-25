package models

type CListsToCTableLink struct {
	CListId  int `db:"cardlist_id"`
	CTableId int `db:"cardtable_id"`
}

type CardsToClistLink struct {
	CListId int `db:"cardlist_id"`
	CardId  int `db:"card_id"`
}
