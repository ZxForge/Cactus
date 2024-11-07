package model

import (
	"database/sql"
	"time"
)

type Pipeline struct {
	Id        int64        `db:"id"`
	TimeStart time.Time    `db:"time_start"`
	TimeEnd   sql.NullTime `db:"time_end"`
}
