package server

import (
	"repo/request"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
	"gorm.io/gorm"
)

type BootApp struct {
	Engine       *gin.Engine
	DB           *gorm.DB
	Oauth        *server.Server
	ClientStore  *store.ClientStore
	OauthManager *manage.Manager
	ConfigAuth   *request.Key
}
