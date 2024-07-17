package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateBookRequest represents the request payload for creating a book
type CreateBookRequest struct {
	Title       string `json:"title" example:"Judul Buku5"`
	Description string `json:"description" example:"Deskripsi buku"`
	ImageURL    string `json:"image_url" example:"https://example.com/gambar.jpg"`
	ReleaseYear int    `json:"release_year" example:"2021"`
	Price       string `json:"price" example:"29.99"`
	TotalPage   int    `json:"total_page" example:"300"`
	UserID      uint   `json:"user_id" example:"1"` // Include user_id in the request payload
}

// CreateBook godoc
// @Summary Create a new book
// @Description Create a new book with the input payload
// @Tags books
// @Accept  json
// @Produce  json
// @Param book body CreateBookRequest true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books [post]
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var bookReq CreateBookRequest

	if err := c.ShouldBindJSON(&bookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the user exists
	var user models.User
	if err := db.First(&user, bookReq.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	book := models.Book{
		Title:       bookReq.Title,
		Description: bookReq.Description,
		ImageURL:    bookReq.ImageURL,
		ReleaseYear: bookReq.ReleaseYear,
		Price:       bookReq.Price,
		TotalPage:   bookReq.TotalPage,
		UserID:      bookReq.UserID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// Validate and set thickness
	if err := validateAndSetThickness(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Attempt to create the book in the database
	if err := db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book in the database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Get all books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// GetBookById godoc
// @Summary Get a book by ID
// @Description Get a book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBookRequest represents the request payload for updating a book
type UpdateBookRequest struct {
	Title       string `json:"title" example:"Judul Buku5"`
	Description string `json:"description" example:"Deskripsi buku"`
	ImageURL    string `json:"image_url" example:"https://example.com/gambar.jpg"`
	ReleaseYear int    `json:"release_year" example:"2021"`
	Price       string `json:"price" example:"29.99"`
	TotalPage   int    `json:"total_page" example:"300"`
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update a book with the input payload
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Param book body UpdateBookRequest true "Book"
// @Success 200 {object} models.Book
// @Failure 400 {object} map[string]string
// @Router /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	var bookReq UpdateBookRequest
	if err := c.ShouldBindJSON(&bookReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book.Title = bookReq.Title
	book.Description = bookReq.Description
	book.ImageURL = bookReq.ImageURL
	book.ReleaseYear = bookReq.ReleaseYear
	book.Price = bookReq.Price
	book.TotalPage = bookReq.TotalPage
	book.UpdatedAt = time.Now()

	// Validate and set thickness
	if err := validateAndSetThickness(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&book)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Delete a book by ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var book models.Book

	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not found!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func validateAndSetThickness(book *models.Book) error {
	if book.TotalPage <= 100 {
		book.Thickness = "tipis"
	} else if book.TotalPage <= 200 {
		book.Thickness = "sedang"
	} else {
		book.Thickness = "tebal"
	}

	return nil
}
