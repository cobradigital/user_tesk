package repositories

import (
	"repo/models"

	"gorm.io/gorm"
)

type UserRepositories struct {
	Db *gorm.DB
}

func (r UserRepositories) GetAll() (users []models.Users, err error) {
	err = r.Db.Find(&users).Error
	return
}

func (r UserRepositories) GetbyID(id string) (user models.Users, err error) {
	err = r.Db.Where("id = ?", id).First(&user).Error
	return
}

func (r UserRepositories) GetByAuthorization(email, password string) (user models.Users, err error) {
	err = r.Db.Where("email = ? AND password = ?", email, password).First(&user).Error
	return
}

func (r UserRepositories) CreateUser(user models.Users) error {
	return r.Db.Create(&user).Error
}

func (r UserRepositories) UpdateUser(user models.Users) error {
	return r.Db.Model(models.Users{}).Where("id = ?", user.ID).Updates(models.Users{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}).Error
}

func (r UserRepositories) DeleteUser(id string) error {
	return r.Db.Where("id = ?", id).Delete(models.Users{}).Error
}
