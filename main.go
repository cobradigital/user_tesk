package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"repo/repositories"
	"repo/request"
	start "repo/server"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

func main() {

	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("err read .env; please change .env.example to .env %v", err)
	}

	c := gin.Default()

	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultAuthorizeCodeTokenCfg)

	clientStore := store.NewClientStore()
	manager.MapClientStorage(clientStore)

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	db, err := repositories.Connect(
		viper.Get("DBHOST").(string),
		viper.Get("DBPORT").(string),
		viper.Get("DBUSER").(string),
		viper.Get("DBPASS").(string),
		viper.Get("DBNAME").(string))
	if err != nil {
		log.Fatalln("err Connect DB : %v", err)
	}

	configAuth := request.Key{
		ClientID:     viper.Get("CLIENTID").(string),
		ClientSecret: viper.Get("CLIENTSECRET").(string),
	}

	boot := start.BootApp{
		Engine:       c,
		DB:           db,
		Oauth:        srv,
		ClientStore:  clientStore,
		OauthManager: manager,
		ConfigAuth:   &configAuth,
	}

	boot.Router()
	serv := &http.Server{
		Addr:    ":" + viper.Get("PORT").(string),
		Handler: boot.Engine,
	}

	go func() {
		// service connections
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
