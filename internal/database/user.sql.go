// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: user.sql

package database

import (
	"context"
)

const addUser = `-- name: AddUser :one
INSERT INTO "User" (
  "firstName", "lastName", "email", "isEmailVerified", "password", "createdAt", "updatedAt"
) VALUES (
  $1, $2, $3, $4, $5, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
)
RETURNING id, "firstName", "lastName", email, "isEmailVerified", password, "createdAt", "updatedAt"
`

type AddUserParams struct {
	FirstName       string
	LastName        string
	Email           string
	IsEmailVerified bool
	Password        string
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) (User, error) {
	row := q.db.QueryRow(ctx, addUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.IsEmailVerified,
		arg.Password,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsEmailVerified,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM "User"
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, "firstName", "lastName", email, "isEmailVerified", password, "createdAt", "updatedAt" FROM "User"
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsEmailVerified,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, "firstName", "lastName", email, "isEmailVerified", password, "createdAt", "updatedAt" FROM "User"
ORDER BY "firstName"
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.Email,
			&i.IsEmailVerified,
			&i.Password,
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

const updateUser = `-- name: UpdateUser :one
UPDATE "User"
SET "firstName" = $1,
    "lastName" = $2,
    "email" = $3,
    "isEmailVerified" = $4,
    "password" = $5,
    "updatedAt" = CURRENT_TIMESTAMP
WHERE id = $6 RETURNING id, "firstName", "lastName", email, "isEmailVerified", password, "createdAt", "updatedAt"
`

type UpdateUserParams struct {
	FirstName       string
	LastName        string
	Email           string
	IsEmailVerified bool
	Password        string
	ID              int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.IsEmailVerified,
		arg.Password,
		arg.ID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.IsEmailVerified,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
