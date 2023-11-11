-- name: CreateMatch :one
INSERT INTO matches (id, hometeam ,awayteam ,datetime ,stadium ,status)  
VALUES ($1, $2, $3, $4, $5, $6)

RETURNING *;

