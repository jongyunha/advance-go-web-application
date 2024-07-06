package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/converter"
	"github.com/jongyunha/advance-go-web-application/api/dto"
	"github.com/jongyunha/advance-go-web-application/api/entity"
)

type UserService interface {
	Create(ctx context.Context, request *dto.CreateUserRequest) (int64, error)
}

type DefaultUserService struct {
	transactionManager entity.TransactionManager
	userRepository     entity.UserRepository
}

func NewDefaultUserService(
	transactionManager entity.TransactionManager,
	userRepository entity.UserRepository,
) *DefaultUserService {
	return &DefaultUserService{
		transactionManager: transactionManager,
		userRepository:     userRepository,
	}
}

func (s *DefaultUserService) Create(ctx context.Context, request *dto.CreateUserRequest) (int64, error) {
	user := converter.ConvertUserRequestToUserEntity(request)
	err := s.transactionManager.Do(func(tx *sqlx.Tx) error {
		return s.userRepository.Create(ctx, tx, user)
	})

	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return *user.ID, nil
}
