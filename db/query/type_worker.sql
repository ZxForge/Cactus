-- name: GetTypeWorkerBySlug :one
SELECT * FROM type_worker
WHERE slug = $1
LIMIT 1;