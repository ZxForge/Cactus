-- name: CreateFile :one
INSERT INTO file (
    id_message, 
    title, 
    "path", 
    ext, 
    "uuid"
) VALUES ($1, $2, $3, $4, $5) 
RETURNING *;