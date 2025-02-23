package controllers

import (
	"books-management-system/internal/models"
	"books-management-system/internal/services"
	"books-management-system/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController struct {
	Service *services.BookService
}

func (c *BookController) InitRoutes(router *gin.Engine) {
	book := router.Group("/books")
	{
		book.GET("", c.GetBooks)
		book.GET("/:id", c.GetBook)
		book.POST("", c.CreateBook)
		book.PUT("/:id", c.UpdateBook)
		book.DELETE("/:id", c.DeleteBook)
	}
}

func NewBookController(service *services.BookService) *BookController {
	return &BookController{Service: service}
}

// GetBooks
// @Summary Get Books
// @Description Fetch paginated list of books
// @Tags books
// @Accept  json
// @Produce  json
// @Param page query int false "Page number"
// @Param limit query int false "Limit per page"
// @Success 200 {array} models.Book
// @Failure 500 {object} gin.H "internal server error"
// @Router /books [get]
func (c *BookController) GetBooks(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	books, err := c.Service.GetBooks(ctx.Request.Context(), page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrInternalError.Error()})
		return
	}
	ctx.JSON(http.StatusOK, books)
}

// GetBook
// @Summary Get Book
// @Description Fetch book details by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 404 {object} gin.H "book not found"
// @Router /books/{id} [get]
func (c *BookController) GetBook(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidBookID.Error()})
		return
	}

	book, err := c.Service.GetBookByID(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": utils.ErrBookNotFound.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// CreateBook
// @Summary Create a new Book
// @Description Add a new book to the system
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body models.Book true "Book data"
// @Success 201 {object} models.Book
// @Failure 400 {object} gin.H "invalid input data"
// @Router /books [post]
func (c *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidInput.Error()})
		return
	}

	if err := utils.ValidateStruct(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.Service.CreateBook(ctx.Request.Context(), &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}
	ctx.JSON(http.StatusCreated, book)
}

// UpdateBook
// @Summary Update a book
// @Description Update an existing book's details
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body models.Book true "Updated book data"
// @Success 200 {object} models.Book
// @Failure 400 {object} gin.H "Invalid input"
// @Failure 500 {object} gin.H "Failed to update book"
// @Router /books/{id} [put]
func (c *BookController) UpdateBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var book models.Book
	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": utils.ErrInvalidInput.Error()})
		return
	}
	book.ID = uint(id)
	err := c.Service.UpdateBook(ctx.Request.Context(), &book)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrBookUpdate.Error()})
		return
	}
	ctx.JSON(http.StatusOK, book)
}

// DeleteBook
// @Summary Delete a book
// @Description Remove a book from the system by its ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} gin.H "Book deleted successfully"
// @Failure 500 {object} gin.H "failed to delete book"
// @Router /books/{id} [delete]
func (c *BookController) DeleteBook(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	err := c.Service.DeleteBook(ctx.Request.Context(), uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": utils.ErrBookDeletion.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
