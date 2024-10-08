// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: matches.sql

package database

import (
	"context"
)

const createMatch = `-- name: CreateMatch :one
INSERT INTO matches (id, hometeam ,awayteam ,datetime ,stadium ,status)  
VALUES ($1, $2, $3, $4, $5, $6)

RETURNING id, hometeam, awayteam, datetime, stadium, status
`

type CreateMatchParams struct {
	ID       string `json:"id"`
	Hometeam string `json:"hometeam"`
	Awayteam string `json:"awayteam"`
	Datetime string `json:"datetime"`
	Stadium  string `json:"stadium"`
	Status   string `json:"status"`
}

func (q *Queries) CreateMatch(ctx context.Context, arg CreateMatchParams) (Match, error) {
	row := q.db.QueryRowContext(ctx, createMatch,
		arg.ID,
		arg.Hometeam,
		arg.Awayteam,
		arg.Datetime,
		arg.Stadium,
		arg.Status,
	)
	var i Match
	err := row.Scan(
		&i.ID,
		&i.Hometeam,
		&i.Awayteam,
		&i.Datetime,
		&i.Stadium,
		&i.Status,
	)
	return i, err
}
