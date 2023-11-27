package handler

import (
	"repo/request"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"gorm.io/gorm"
)

type Handler struct {
	Db          *gorm.DB
	Oauth       *server.Server
	ClientStore *store.ClientStore
	Manager     *manage.Manager
	ConfigAuth  *request.Key
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func response(c *gin.Context, code int, result interface{}) {
	c.JSON(code, Response{
		Code: code,
		Data: result,
	})
}
