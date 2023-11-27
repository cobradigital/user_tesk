package handler

import (
	"fmt"
	"log"
	"net/http"
	"repo/request"
	"repo/services"

	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
)

type AuthHandler struct {
	Handler
}

func (h AuthHandler) Token(c *gin.Context) {
	h.Oauth.HandleTokenRequest(c.Writer, c.Request)
}

func (h AuthHandler) Credential(c *gin.Context) {
	auth := request.Auth{
		Scope:        c.Query("scope"),
		GrantType:    c.Query("grant_type"),
		ClientID:     h.ConfigAuth.ClientID,
		ClientSecret: h.ConfigAuth.ClientSecret,
		Username:     c.Query("username"),
		Password:     c.Query("password"),
	}
	authService := services.AuthService{Db: h.Db}
	_, data, err := authService.GetByAuthorization(auth.Username, auth.Password)

	h.Manager.MustTokenStorage(store.NewMemoryTokenStore())

	h.Oauth.SetAllowGetAccessRequest(true)
	h.Oauth.SetClientInfoHandler(h.Oauth.ClientInfoHandler)
	err = h.ClientStore.Set(auth.ClientID, &models.Client{
		ID:     auth.ClientID,
		Secret: auth.ClientSecret,
		Domain: "http://localhost:8080",
		UserID: fmt.Sprintf("%d", data.ID),
	})
	if err != nil {
		log.Println(err.Error())
	}

	c.JSON(http.StatusOK, map[string]string{"USER_ID": fmt.Sprintf("%d", data.ID), "CLIENT_ID": auth.ClientID, "CLIENT_SECRET": auth.ClientSecret})
}
