-- name: CreateUser :one
INSERT INTO "user" (
    fio, 
    login, 
    email, 
    "password", 
    reset_password_after_login, 
    create_at
) VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;