package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temaxuck/WUR/service.ebooks/pkg/models"
)

type CreateBookRequestBody struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	AuthorID        uint   `json:"author_id"`
	// Cover	string
}

func (h Handler) CreateBook(ctx *gin.Context) {
	body := CreateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var bookMeta models.BookMeta

	bookMeta.Title = body.Title
	bookMeta.Description = body.Description
	pubDate, _ := time.Parse("02.01.2006", body.PublicationDate)
	bookMeta.PublicationDate = pubDate

	if body.AuthorID == 0 {
		author, err := fetchDefaultAuthor(h.DB)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		bookMeta.AuthorID = author.ID
	} else {
		bookMeta.AuthorID = body.AuthorID
	}

	var author models.Author

	if result := h.DB.First(&author, bookMeta.AuthorID); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	bookMeta.Author = author

	if result := h.DB.Create(&bookMeta); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &bookMeta)
}

type UpdateBookRequestBody struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	AuthorID        uint   `json:"author_id"`
	// Cover	string
}

func (h Handler) UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateBookRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var bookMeta models.BookMeta

	if result := h.DB.First(&bookMeta, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	bookMeta.Title = body.Title
	bookMeta.Description = body.Description
	pubDate, _ := time.Parse("02.01.2006", body.PublicationDate)
	bookMeta.PublicationDate = pubDate

	if body.AuthorID == 0 {
		author, err := fetchDefaultAuthor(h.DB)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		bookMeta.AuthorID = author.ID
	} else {
		bookMeta.AuthorID = body.AuthorID
	}

	var author models.Author

	if result := h.DB.First(&author, bookMeta.AuthorID); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	bookMeta.Author = author

	h.DB.Save(&bookMeta)

	ctx.JSON(http.StatusOK, &bookMeta)
}

func (h Handler) GetBooks(ctx *gin.Context) {
	var bookMetas []models.BookMeta

	if result := h.DB.Preload("Author").Find(&bookMetas); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &bookMetas)
}

func (h Handler) GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var bookMeta models.BookMeta

	if result := h.DB.Preload("Author").First(&bookMeta, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &bookMeta)
}

func (h Handler) DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")

	var bookMeta models.BookMeta

	if result := h.DB.First(&bookMeta, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	h.DB.Delete(&bookMeta)

	ctx.Status(http.StatusOK)
}
