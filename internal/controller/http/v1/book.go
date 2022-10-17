package v1

import (
	"net/http"
	"pet-project-1/internal/entity"
	"pet-project-1/internal/usecase"

	"github.com/gin-gonic/gin"
)

type bookRoutes struct {
	b usecase.Book
}

func newBookRoutes(handler *gin.RouterGroup, bk usecase.Book) {
	br := &bookRoutes{b:bk}

	handler.GET("/books", br.getBooks)

}

type bookResponse struct {
	Books []entity.Book `json:"books"`
}

func (br *bookRoutes) getBooks(c *gin.Context) {
	listBooks, err := br.b.GetBooks(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	}
	c.JSON(http.StatusOK, bookResponse{listBooks})
}
