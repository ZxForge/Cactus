-- name: GetProcessBySlug :one
SELECT * FROM process
WHERE slug = $1
LIMIT 1;