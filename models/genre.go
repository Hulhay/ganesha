package models

type Genre struct {
	ID   int64  `db:"genre_id"`
	Uuid string `db:"genre_uuid"`
	Name string `db:"genre_name"`
}
