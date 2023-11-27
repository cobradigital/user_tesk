package server

import (
	"repo/handler"
	"repo/middleware"

	"github.com/gin-gonic/gin"
)

func (boot *BootApp) Router() {

	setup := handler.Handler{
		Db:          boot.DB,
		Oauth:       boot.Oauth,
		ClientStore: boot.ClientStore,
		Manager:     boot.OauthManager,
		ConfigAuth:  boot.ConfigAuth,
	}

	boot.Engine.Use(gin.Logger())
	boot.Engine.Use(gin.Recovery())

	userHandler := handler.UserHandler{setup}
	user := boot.Engine.Group("/users")
	{
		user.GET("/", userHandler.GetAll)
		user.GET("/:id", userHandler.GetByID)
		user.POST("/", userHandler.PostUser)
		user.PUT("/:id", userHandler.PutUser)
		user.DELETE("/:id", userHandler.DeleteUser)
	}

	authHandler := handler.AuthHandler{setup}
	boot.Engine.POST("/token", authHandler.Token)
	boot.Engine.GET("/credentials", authHandler.Credential)

	taskHandler := handler.TaskHandler{setup}
	task := boot.Engine.Group("/task")
	task.Use(middleware.Auth(setup.Oauth))
	{
		task.GET("/", taskHandler.GetAll)
		task.GET("/:id", taskHandler.GetByID)
		task.POST("/", taskHandler.PostTask)
		task.PUT("/:id", taskHandler.PutTask)
		task.DELETE("/:id", taskHandler.DeleteTask)
	}

}
