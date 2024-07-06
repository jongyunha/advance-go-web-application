package service

import "github.com/jongyunha/advance-go-web-application/api/entity"

type UserService interface {
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
