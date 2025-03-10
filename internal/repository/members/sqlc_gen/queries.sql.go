// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: queries.sql

package sqlc_members

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createMember = `-- name: CreateMember :one
INSERT INTO members.members_table(
    name
    , birthday
) VALUES(
    $1, $2
) RETURNING id, name, birthday
`

type CreateMemberParams struct {
	Name     string
	Birthday pgtype.Date
}

func (q *Queries) CreateMember(ctx context.Context, db DBTX, arg *CreateMemberParams) (*MembersMembersTable, error) {
	row := db.QueryRow(ctx, createMember, arg.Name, arg.Birthday)
	var i MembersMembersTable
	err := row.Scan(&i.ID, &i.Name, &i.Birthday)
	return &i, err
}

const deleteMember = `-- name: DeleteMember :one
DELETE FROM members.members_table
WHERE
    id = $1
RETURNING id, name, birthday
`

func (q *Queries) DeleteMember(ctx context.Context, db DBTX, id int64) (*MembersMembersTable, error) {
	row := db.QueryRow(ctx, deleteMember, id)
	var i MembersMembersTable
	err := row.Scan(&i.ID, &i.Name, &i.Birthday)
	return &i, err
}

const getMember = `-- name: GetMember :one
SELECT id, name, birthday
FROM members.members_table
WHERE
    id = $1
LIMIT 1
`

func (q *Queries) GetMember(ctx context.Context, db DBTX, id int64) (*MembersMembersTable, error) {
	row := db.QueryRow(ctx, getMember, id)
	var i MembersMembersTable
	err := row.Scan(&i.ID, &i.Name, &i.Birthday)
	return &i, err
}

const getMembersList = `-- name: GetMembersList :many
SELECT id, name, birthday
FROM members.members_table mt
WHERE
    $1::text is null
    OR
    mt.name ILIKE('%'||$1::text||'%')
ORDER BY
    mt.name
`

func (q *Queries) GetMembersList(ctx context.Context, db DBTX, name pgtype.Text) ([]*MembersMembersTable, error) {
	rows, err := db.Query(ctx, getMembersList, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []*MembersMembersTable
	for rows.Next() {
		var i MembersMembersTable
		if err := rows.Scan(&i.ID, &i.Name, &i.Birthday); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMember = `-- name: UpdateMember :one
UPDATE members.members_table
SET
    name = $1
    , birthday = $2
WHERE
    id = $3
RETURNING id, name, birthday
`

type UpdateMemberParams struct {
	Name     string
	Birthday pgtype.Date
	ID       int64
}

func (q *Queries) UpdateMember(ctx context.Context, db DBTX, arg *UpdateMemberParams) (*MembersMembersTable, error) {
	row := db.QueryRow(ctx, updateMember, arg.Name, arg.Birthday, arg.ID)
	var i MembersMembersTable
	err := row.Scan(&i.ID, &i.Name, &i.Birthday)
	return &i, err
}
