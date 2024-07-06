package service

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/jongyunha/advance-go-web-application/api/converter"
	"github.com/jongyunha/advance-go-web-application/api/dto"
	"github.com/jongyunha/advance-go-web-application/api/entity"
	"github.com/jongyunha/advance-go-web-application/api/errs"
	"net/http"
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
	exists, err := s.userRepository.ExistsByEmail(ctx, request.Email)
	if err != nil {
		return 0, errs.NewError(http.StatusInternalServerError, errs.InternalServerError, fmt.Sprintf("failed to check email existence: %v", err))
	} else if exists {
		return 0, errs.NewWarnError(http.StatusBadRequest, errs.InvalidRequest, fmt.Sprintf("email %s already exists", request.Email))
	}

	user := converter.ConvertUserRequestToUserEntity(request)
	err = user.HashPassword(request.Password)
	if err != nil {
		return 0, errs.NewError(http.StatusInternalServerError, errs.InternalServerError, fmt.Sprintf("failed to hash password: %v", err))
	}

	if err := s.transactionManager.Do(func(tx *sqlx.Tx) error {
		return s.userRepository.Create(ctx, tx, user)
	}); err != nil {
		return 0, errs.NewError(http.StatusInternalServerError, errs.InternalServerError, fmt.Sprintf("failed to create user: %v", err))
	}

	return *user.ID, nil
}
