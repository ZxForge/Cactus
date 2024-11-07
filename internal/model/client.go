package model

type Client struct {
	Id       int64  `db:"id"`
	Name     string `db:"name"`
	Token    string `db:"token"`
	Action   bool   `db:"action"`
	Priority uint8  `db:"priority"`
}
