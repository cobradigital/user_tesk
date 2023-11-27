package handler

import (
	"net/http"
	"repo/request"
	"repo/services"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	Handler
}

func (h TaskHandler) GetAll(c *gin.Context) {

	taskService := services.TaskService{Db: h.Db}
	code, data, err := taskService.GetAll()
	if err != nil {
		response(c, http.StatusNotFound, "Not Found")
		return
	}

	response(c, code, data)
}

func (h TaskHandler) GetByID(c *gin.Context) {

	var request request.ByID

	if err := c.ShouldBindUri(&request); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}
	taskService := services.TaskService{Db: h.Db}
	code, data, err := taskService.GetByID(request.ID)
	if err != nil {
		response(c, http.StatusNotFound, "Not Found")
		return
	}

	response(c, code, data)
}

func (h TaskHandler) PostTask(c *gin.Context) {
	var task request.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	taskService := services.TaskService{Db: h.Db}
	code, msg, err := taskService.Create(task)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}

func (h TaskHandler) PutTask(c *gin.Context) {
	var id request.ByID
	var task request.Task

	if err := c.ShouldBindUri(&id); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		response(c, http.StatusBadRequest, "Bad Request Data")
		return
	}

	task.ID = id.ID
	taskService := services.TaskService{Db: h.Db}
	code, msg, err := taskService.Update(task)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}

func (h TaskHandler) DeleteTask(c *gin.Context) {

	var request request.ByID

	if err := c.ShouldBindUri(&request); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}
	taskService := services.TaskService{Db: h.Db}
	code, msg, err := taskService.Delete(request.ID)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}
