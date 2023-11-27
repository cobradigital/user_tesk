package services

import (
	"net/http"
	"repo/models"
	"repo/repositories"
	"repo/request"
	"repo/services/views"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type UserService struct {
	Db *gorm.DB
}

func (s UserService) GetAll() (code int, result []views.User, err error) {
	userRepo := repositories.UserRepositories{Db: s.Db}
	data, err := userRepo.GetAll()
	if err != nil {
		return http.StatusNotFound, result, err
	}

	for _, d := range data {
		result = append(result, views.User{
			ID:        d.ID,
			Name:      d.Name,
			Email:     d.Email,
			CreatedAt: d.CreatedAt,
		})
	}

	return http.StatusOK, result, nil
}

func (s UserService) GetByID(id string) (code int, result views.User, err error) {
	userRepo := repositories.UserRepositories{Db: s.Db}
	data, err := userRepo.GetbyID(id)
	if err != nil {
		return http.StatusNotFound, result, err
	}

	return http.StatusOK, views.User{
		ID:        data.ID,
		Name:      data.Name,
		Email:     data.Email,
		CreatedAt: data.CreatedAt,
	}, nil
}

func (s UserService) Create(user request.User) (int, string, error) {

	userRepo := repositories.UserRepositories{Db: s.Db}
	if err := userRepo.CreateUser(models.Users{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
	}); err != nil {
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success Create User", nil
}

func (s UserService) Update(user request.UpdateUser) (int, string, error) {

	id, _ := strconv.Atoi(user.ID)
	userRepo := repositories.UserRepositories{Db: s.Db}
	if err := userRepo.UpdateUser(models.Users{
		ID:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}); err != nil {
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success update User", nil
}

func (s UserService) Delete(id string) (int, string, error) {

	userRepo := repositories.UserRepositories{Db: s.Db}
	if err := userRepo.DeleteUser(id); err != nil {
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success hapus User", nil
}
