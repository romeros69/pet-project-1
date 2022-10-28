package v1

import (
	"pet-project-1/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(handler *gin.Engine, bk usecase.Book) {
	h := handler.Group("/api/v1")
	{
		newBookRoutes(h, bk)
	}
}
