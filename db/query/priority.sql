-- name: GetPriorityBySlug :one
SELECT * FROM priority p
WHERE p.slug = $1
LIMIT 1;

-- name: GetPriorityBySystemId :one
SELECT * FROM priority p 
JOIN system_client sc ON sc.id_priority = p.id
WHERE sc.id = $1;