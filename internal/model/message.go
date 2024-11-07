package model

import (
	"database/sql"
	"time"
)

type Message struct {
	Id int64 `db:"id"`

	ProcessId int64 `db:"process_id"`
	ClientId  int64 `db:"client_id"`

	UUID      string       `db:"UUID"`
	Title     string       `db:"title"`
	Value     string       `db:"value"`
	Priority  uint8        `db:"priority"` // TODO в базе int32
	SendLater sql.NullTime `db:"send_later"`
	CreateAT  time.Time    `db:"created_at"`
}
