package model

type Role struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
	Slug string `db:"slug"`
}
