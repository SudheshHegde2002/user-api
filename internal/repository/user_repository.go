package repository

import (
	"context"
	"time"
	"user-api/db/sqlc"

	"github.com/jackc/pgx/v5/pgtype"
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
	parsedDob, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, err
	}
	return r.queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: name,
		Dob:  pgtype.Date{Time: parsedDob, Valid: true},
	})
}

func (r *UserRepository) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, error) {
	return r.queries.GetUserByID(ctx, id)
}
