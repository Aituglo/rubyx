// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: ip.sql

package db

import (
	"context"
)

const createIp = `-- name: CreateIp :one
INSERT INTO ip (program_id, subdomain_id, ip) VALUES ($1, $2, $3) RETURNING id, program_id, subdomain_id, ip, created_at, updated_at
`

type CreateIpParams struct {
	ProgramID   int64  `json:"program_id"`
	SubdomainID int64  `json:"subdomain_id"`
	Ip          string `json:"ip"`
}

func (q *Queries) CreateIp(ctx context.Context, arg CreateIpParams) (Ip, error) {
	row := q.db.QueryRow(ctx, createIp, arg.ProgramID, arg.SubdomainID, arg.Ip)
	var i Ip
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.SubdomainID,
		&i.Ip,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteIpByIDs = `-- name: DeleteIpByIDs :exec
DELETE FROM ip WHERE id = $1
`

func (q *Queries) DeleteIpByIDs(ctx context.Context, id int64) error {
	_, err := q.db.Exec(ctx, deleteIpByIDs, id)
	return err
}

const findIpByIDs = `-- name: FindIpByIDs :one
SELECT id, program_id, subdomain_id, ip, created_at, updated_at FROM ip WHERE id = $1 LIMIT 1
`

func (q *Queries) FindIpByIDs(ctx context.Context, id int64) (Ip, error) {
	row := q.db.QueryRow(ctx, findIpByIDs, id)
	var i Ip
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.SubdomainID,
		&i.Ip,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const findIps = `-- name: FindIps :many
SELECT id, program_id, subdomain_id, ip, created_at, updated_at FROM ip
`

func (q *Queries) FindIps(ctx context.Context) ([]Ip, error) {
	rows, err := q.db.Query(ctx, findIps)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Ip{}
	for rows.Next() {
		var i Ip
		if err := rows.Scan(
			&i.ID,
			&i.ProgramID,
			&i.SubdomainID,
			&i.Ip,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateIp = `-- name: UpdateIp :one
UPDATE ip SET program_id = $2, subdomain_id = $3, ip = $4, updated_at = NOW() WHERE id = $1 RETURNING id, program_id, subdomain_id, ip, created_at, updated_at
`

type UpdateIpParams struct {
	ID          int64  `json:"id"`
	ProgramID   int64  `json:"program_id"`
	SubdomainID int64  `json:"subdomain_id"`
	Ip          string `json:"ip"`
}

func (q *Queries) UpdateIp(ctx context.Context, arg UpdateIpParams) (Ip, error) {
	row := q.db.QueryRow(ctx, updateIp,
		arg.ID,
		arg.ProgramID,
		arg.SubdomainID,
		arg.Ip,
	)
	var i Ip
	err := row.Scan(
		&i.ID,
		&i.ProgramID,
		&i.SubdomainID,
		&i.Ip,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
