package model

import "time"

type User struct {
	Id                      int64     `db:"id"`
	Login                   string    `db:"login"`
	Email                   string    `db:"email"`
	Password                string    `db:"password"`
	Role                    int64     `db:"role_id"`
	ResetPasswordAfterLogin bool      `db:"reset_password_after_login"`
	CreateAT                time.Time `db:"created_at"`
}
