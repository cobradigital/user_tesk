package repositories

import (
	"repo/models"

	"gorm.io/gorm"
)

type TaskRepositories struct {
	Db *gorm.DB
}

func (r TaskRepositories) GetAll() (tasks []models.Tasks, err error) {
	err = r.Db.Find(&tasks).Error
	return
}

func (r TaskRepositories) GetbyID(id string) (task models.Tasks, err error) {
	err = r.Db.Where("id = ?", id).First(&task).Error
	return
}

func (r TaskRepositories) Createtask(task models.Tasks) error {
	return r.Db.Create(&task).Error
}

func (r TaskRepositories) Updatetask(task models.Tasks) error {
	return r.Db.Model(models.Tasks{}).Where("id = ?", task.ID).Updates(models.Tasks{
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}).Error
}

func (r TaskRepositories) Deletetask(id string) error {
	return r.Db.Where("id = ?", id).Delete(models.Tasks{}).Error
}
