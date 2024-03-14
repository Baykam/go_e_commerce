package service

import (
	"golang_testing_grpc/internal/user/repository"

	"github.com/quangdangfit/gocommon/validation"
)

type IUserService interface{}

type UserService struct {
	validation validation.Validation
	repo       repository.IUserRepository
}

func NewUserService(validation validation.Validation, repo repository.IUserRepository) *UserService {
	return &UserService{
		validation: validation,
		repo:       repo,
	}
}

// func (s *UserService) GetUserAll(ctx context.Context) ([]*model.User, error) {
// 	list, err := s.repo.GetAllUsers(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return list, nil
// }
