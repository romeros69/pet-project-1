package app

import (
	"log"
	"os"
	"os/signal"
	"pet-project-1/config"
	v1 "pet-project-1/internal/controller/http/v1"
	"pet-project-1/internal/usecase"
	"pet-project-1/internal/usecase/repo"
	"pet-project-1/pkg/httpserver"
	"pet-project-1/pkg/postgres"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(cfg *config.Config) {
	pg, err := postgres.New(cfg)

	if err != nil {
		log.Fatal("Error in creating postgres instance")
	}

	bookUseCase := usecase.NewBookUseCase(repo.NewBookRepo(pg))

	// http server
	handler := gin.New()

	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Access-Control-Allow-Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	v1.NewRouter(handler, bookUseCase)

	serv := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))
	interruption := make(chan os.Signal, 1)
	signal.Notify(interruption, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interruption:
		log.Printf("signal: " + s.String())
	case err = <-serv.Notify():
		log.Printf("Notify from http server")
	}

	err = serv.Shutdown()
	if err != nil {
		log.Printf("Http server shutdown")
	}
}
