-- name: GetChirp :one
SELECT * FROM chirps_table
WHERE id = $1;