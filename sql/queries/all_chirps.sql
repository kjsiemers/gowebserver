-- name: GetAllChirps :many
SELECT id, created_at, updated_at, body, user_id
FROM chirps_table
ORDER BY created_at;
