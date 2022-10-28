package v1

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"pet-project-1/internal/usecase"

	"github.com/gin-gonic/gin"
	_ "pet-project-1/docs"
)

func NewRouter(handler *gin.Engine, bk usecase.Book) {
	handler.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	h := handler.Group("/api/v1")
	{
		newBookRoutes(h, bk)
	}
}
