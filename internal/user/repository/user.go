package repository

import (
	"context"
	"golang_testing_grpc/internal/user/model"
	"golang_testing_grpc/pkg/db"
)

type UserRepo struct {
	db db.IDatabaseInterface
}

func NewUserRepo(db db.IDatabaseInterface) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	_, err := r.db.GetFromTableWithConditionInList(ctx, db.UsersTable, db.Id, user)
	if err != nil {
		return err
	}
	return nil
}
