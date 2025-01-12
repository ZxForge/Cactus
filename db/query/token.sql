-- name: CreateToken :one
INSERT INTO "token" (
    id_system, 
    id_kind_worker, 
    is_active, 
    public_token, 
    secret_token
) VALUES ($1, $2, $3, $4, $5)
RETURNING *;