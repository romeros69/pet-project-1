package v1

import (
	"fmt"
	"net/http"
	"pet-project-1/internal/entity"
	"pet-project-1/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type bookRoutes struct {
	b usecase.Book
}

func newBookRoutes(handler *gin.RouterGroup, bk usecase.Book) {
	br := &bookRoutes{b: bk}

	handler.GET("/book", br.getBooks)
	handler.POST("/book", br.createBook)
	handler.DELETE("/book/:id", br.deleteBook)
	handler.PUT("/book/:id", br.updateBook)

}

// ok
func (br *bookRoutes) getBooks(c *gin.Context) {
	listBooks, err := br.b.GetBooks(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
	}
	var req []entity.Book
	if listBooks != nil {
		req = listBooks
	} else {
		req = make([]entity.Book, 0)
	}
	c.JSON(http.StatusOK, req)
}

type bookRequest struct {
	Tittle string `json:"tittle"`
	Author string `json:"author"`
}

//ok
func (br *bookRoutes) createBook(c *gin.Context) {
	req := new(bookRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}
	id, err := br.b.CreateBook(c.Request.Context(), entity.Book{
		Tittle: req.Tittle,
		Author: req.Author,
	})
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Header("Location", fmt.Sprintf("/api/v1/book/%s", id.String()))
	c.JSON(http.StatusCreated, nil)
}

func (br *bookRoutes) deleteBook(c *gin.Context) {
	ID := c.Param("id")
	if ID == "" {
		errorResponse(c, http.StatusBadRequest, "Отсутствует параметр id")
	}
	err := br.b.DeleteBook(c.Request.Context(), ID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

func (br *bookRoutes) updateBook(c *gin.Context) {
	bookID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}

	req := new(bookRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
	}

	err = br.b.UpdateBook(c.Request.Context(), entity.Book{
		ID:     bookID,
		Tittle: req.Tittle,
		Author: req.Author,
	})

	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
