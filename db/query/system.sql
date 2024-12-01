-- name: GetSystemIdByToken :one
SELECT sc.id
FROM system_client sc
INNER JOIN "token" t ON t.id_system = sc.id
WHERE token = $1
LIMIT 1;

