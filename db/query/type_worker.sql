-- name: GetTypeWorkerBySlug :one
SELECT * FROM type_worker
WHERE slug = $1
LIMIT 1;

-- name: GetTypeWorkers :many
SELECT * FROM type_worker;