package services

import (
	"log"
	"net/http"
	"repo/models"
	"repo/repositories"
	"repo/request"
	"repo/services/views"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type TaskService struct {
	Db *gorm.DB
}

func (s TaskService) GetAll() (code int, result []views.Tasks, err error) {
	taskRepo := repositories.TaskRepositories{Db: s.Db}
	data, err := taskRepo.GetAll()
	if err != nil {
		return http.StatusNotFound, result, err
	}

	for _, d := range data {
		result = append(result, views.Tasks{
			ID:          d.ID,
			UserID:      d.UserID,
			Title:       d.Title,
			Description: d.Description,
			Status:      d.Status,
			CreatedAt:   d.CreatedAt,
			UpdatedAt:   d.UpdatedAt,
		})
	}

	return http.StatusOK, result, nil
}

func (s TaskService) GetByID(id string) (code int, result views.Tasks, err error) {
	taskRepo := repositories.TaskRepositories{Db: s.Db}
	data, err := taskRepo.GetbyID(id)
	if err != nil {
		return http.StatusNotFound, result, err
	}

	return http.StatusOK, views.Tasks{
		ID:          data.ID,
		UserID:      data.UserID,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}, nil
}

func (s TaskService) Create(task request.Task) (int, string, error) {

	taskRepo := repositories.TaskRepositories{Db: s.Db}
	if err := taskRepo.Createtask(models.Tasks{
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   time.Now(),
	}); err != nil {
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success Create Task", nil
}

func (s TaskService) Update(task request.Task) (int, string, error) {

	id, _ := strconv.Atoi(task.ID)
	taskRepo := repositories.TaskRepositories{Db: s.Db}
	if err := taskRepo.Updatetask(models.Tasks{
		ID:          id,
		UserID:      task.UserID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
	}); err != nil {
		log.Println(err)
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success update Task", nil
}

func (s TaskService) Delete(id string) (int, string, error) {

	taskRepo := repositories.TaskRepositories{Db: s.Db}
	if err := taskRepo.Deletetask(id); err != nil {
		return http.StatusInternalServerError, "Kesalahan pada system", err
	}

	return http.StatusOK, "Success hapus Task", nil
}
