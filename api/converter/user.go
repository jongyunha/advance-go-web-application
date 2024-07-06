package converter

import (
	"github.com/jongyunha/advance-go-web-application/api/dto"
	"github.com/jongyunha/advance-go-web-application/api/entity"
)

func ConvertUserRequestToUserEntity(request *dto.CreateUserRequest) *entity.User {
	return &entity.User{
		Email:    request.Email,
		UserName: request.UserName,
	}
}
