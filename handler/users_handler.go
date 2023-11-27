package handler

import (
	"net/http"
	"repo/request"
	"repo/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Handler
}

func (h UserHandler) GetAll(c *gin.Context) {

	userService := services.UserService{Db: h.Db}
	code, data, err := userService.GetAll()
	if err != nil {
		response(c, http.StatusNotFound, "Not Found")
		return
	}

	response(c, code, data)
}

func (h UserHandler) GetByID(c *gin.Context) {

	var request request.ByID

	if err := c.ShouldBindUri(&request); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}
	userService := services.UserService{Db: h.Db}
	code, data, err := userService.GetByID(request.ID)
	if err != nil {
		response(c, http.StatusNotFound, "Not Found")
		return
	}

	response(c, code, data)
}

func (h UserHandler) PostUser(c *gin.Context) {
	var user request.User

	if err := c.ShouldBindJSON(&user); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	userService := services.UserService{Db: h.Db}
	code, msg, err := userService.Create(user)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}

func (h UserHandler) PutUser(c *gin.Context) {
	var id request.ByID
	var user request.UpdateUser

	if err := c.ShouldBindUri(&id); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		response(c, http.StatusBadRequest, "Bad Request Data")
		return
	}

	user.ID = id.ID
	userService := services.UserService{Db: h.Db}
	code, msg, err := userService.Update(user)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}

func (h UserHandler) DeleteUser(c *gin.Context) {

	var request request.ByID

	if err := c.ShouldBindUri(&request); err != nil {
		response(c, http.StatusBadRequest, "Bad Request")
		return
	}
	userService := services.UserService{Db: h.Db}
	code, msg, err := userService.Delete(request.ID)
	if err != nil {
		response(c, code, msg)
		return
	}

	response(c, code, msg)
}
