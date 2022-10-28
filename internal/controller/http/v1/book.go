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
	handler.GET("book/:id", br.getBookById)
	handler.POST("/book", br.createBook)
	handler.DELETE("/book/:id", br.deleteBook)
	handler.PUT("/book/:id", br.updateBook)
}

// @Summary GetAllBooks
// @Tags book
// @Description Get all books
// @ID get-all-books
// @Accept json
// @Produce json
// @Success 200 {object} []entity.Book
// @Failure 500 {object} errResponse
// @Router /api/v1/book [get]
func (br *bookRoutes) getBooks(c *gin.Context) {
	listBooks, err := br.b.GetBooks(c.Request.Context())
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var req []entity.Book
	if listBooks != nil {
		req = listBooks
	} else {
		req = make([]entity.Book, 0)
	}
	c.JSON(http.StatusOK, req)
}

// @Summary GetBookById
// @Tags book
// @Description Get book by id
// @ID get-book-by-id
// @Accept json
// @Produce json
// @Param id path string true "Enter id book"
// @Success 200 {object} entity.Book
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Router /api/v1/book/{id} [get]
func (br *bookRoutes) getBookById(c *gin.Context) {
	bookID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	book, err := br.b.GetBookById(c.Request.Context(), bookID)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, book)
}

type bookRequest struct {
	Tittle string `json:"tittle"`
	Author string `json:"author"`
}

// @Summary CreateBook
// @Tags book
// @Description Create new book
// @ID create-book
// @Accept json
// @Produce json
// @Param input body bookRequest true "Enter tittle and author of new book"
// @Success 201 {object} nil
// @Failure 400 {object} errResponse
// @Failure 500 {object} errResponse
// @Router /api/v1/book [post]
func (br *bookRoutes) createBook(c *gin.Context) {
	req := new(bookRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
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

// @Summary DeleteBook
// @Tags book
// @Description Delete book by id
// @ID delete-book
// @Accept json
// @Produce json
// @Param id path string true "Enter id book"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Router /api/v1/book/{id} [delete]
func (br *bookRoutes) deleteBook(c *gin.Context) {
	bookID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "Отсутствует параметр id")
		return
	}
	err = br.b.DeleteBook(c.Request.Context(), bookID)
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary UpdateBook
// @Tags book
// @Description Update book by id
// @ID update-book
// @Accept json
// @Produce json
// @Param id path string true "Enter id book"
// @Param input body bookRequest true "Enter tittle and author of new book"
// @Success 204 {object} nil
// @Failure 400 {object} errResponse
// @Failure 404 {object} errResponse
// @Router /api/v1/book/{id} [put]
func (br *bookRoutes) updateBook(c *gin.Context) {
	bookID, err := uuid.FromString(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	req := new(bookRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = br.b.UpdateBook(c.Request.Context(), entity.Book{
		ID:     bookID,
		Tittle: req.Tittle,
		Author: req.Author,
	})
	if err != nil {
		errorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
