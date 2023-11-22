// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: users.sql

package sqlc

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(name, email, password)
VALUES ($1, $2, $3)
RETURNING id, name, email, password, created_at
`

type CreateUserParams struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Name, arg.Email, arg.Password)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const getUserWithEmail = `-- name: GetUserWithEmail :one
SELECT id, name, email, password, created_at
FROM users
WHERE email = $1
`

func (q *Queries) GetUserWithEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserWithEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
