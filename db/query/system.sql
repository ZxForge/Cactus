-- name: GetSystemIdByToken :one
SELECT s.id
FROM system s
INNER JOIN "token" t ON t.id_system = s.id
WHERE public_token = $1
LIMIT 1;

-- name: CreateSystem :one
INSERT INTO "system" (
    create_user, 
    id_priority, 
    "name", 
    description, 
    is_active
) 
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: AddKindWorkerForSystem :exec
INSERT INTO kind_worker_system (
    id_system, 
    id_kind_worker
) VALUES($1, $2)
RETURNING *;
