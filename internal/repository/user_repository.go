package repository

import (
	"context"
	"user-api/db/sqlc"
)

type UserRepository struct {
	queries *sqlc.Queries
}

func NewUserRepository(queries *sqlc.Queries) *UserRepository {
	return &UserRepository{queries: queries}
}

func (r *UserRepository) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  dob,
	})
}