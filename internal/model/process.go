package model

type Process struct {
	Id   int64  `db:"id"`
	Name string `db:"name"`
}
