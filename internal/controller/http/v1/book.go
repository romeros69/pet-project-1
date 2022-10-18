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
	br := &bookRoutes{b: bk}

	handler.GET("/book", br.getBooks)
	handler.POST("/book", br.createBook)

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

type bookRequest struct {
	Tittle string `json:"tittle"`
	Author string `json:"author"`
}

func (br *bookRoutes) createBook(c *gin.Context) {
	req := new(bookRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}
	err := br.b.CreateBook(c.Request.Context(), entity.Book{
		Tittle: req.Tittle,
		Author: req.Author,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, nil)
}
