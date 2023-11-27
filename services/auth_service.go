package services

import (
	"net/http"
	"repo/repositories"
	"repo/services/views"

	"gorm.io/gorm"
)

type AuthService struct {
	Db *gorm.DB
}

func (s AuthService) GetByAuthorization(email, password string) (code int, result views.User, err error) {
	userRepo := repositories.UserRepositories{Db: s.Db}
	data, err := userRepo.GetByAuthorization(email, password)
	if err != nil {
		return http.StatusNotFound, result, err
	}

	result = views.User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}

	return http.StatusOK, result, nil
}
