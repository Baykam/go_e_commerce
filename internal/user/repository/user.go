package repository

import (
	"context"
	"golang_testing_grpc/internal/user/model"
	"golang_testing_grpc/pkg/db"
)

type IUserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetAllUsers(ctx context.Context) ([]*model.User, error)
}

type UserRepo struct {
	db db.IDatabaseInterface
}

func NewUserRepo(db db.IDatabaseInterface) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	// _, err := r.db.GetFromTableWithConditionInList(ctx, db.UsersTable, db.Id, user)
	// if err != nil {
	// 	return err
	// }
	return nil
}

func (r *UserRepo) GetAllUsers(ctx context.Context) ([]*model.User, error) {
	// var userList []*model.User
	// list, err := r.db.GetAllListFromTable(ctx, db.UsersTable, model.User{})

	// for _, v := range list {
	// 	if user, ok := v.(*model.User); ok {
	// 		userList = append(userList, user)
	// 	}
	// }
	// if err != nil {
	// 	return nil, err
	// }

	return nil, nil
}
