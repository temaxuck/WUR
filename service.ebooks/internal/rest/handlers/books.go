package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
	"github.com/temaxuck/WUR/service.ebooks/internal/exceptions"
	"github.com/temaxuck/WUR/service.ebooks/pkg/models"
)

type CreateBookRequestBody struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	AuthorID        uint   `json:"author_id"`
	// Genres
	// Tags
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

type UploadBookFileRequestBody struct {
	File *multipart.FileHeader `json:"file"`
	// Image	string
}

func (h Handler) UploadBookFile(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UploadBookFileRequestBody{}

	if err := ctx.ShouldBind(&body); err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var bookMeta models.BookMeta
	if result := h.DB.First(&bookMeta, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
	}

	fileHeader, getFileErr := ctx.FormFile("file")
	if getFileErr != nil {
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Couldn't load file."},
		)
		return
	}

	switch ValidateFile(fileHeader, constants.GetBookFileFormats()).(type) {
	case exceptions.FileTooLargeError:
		ctx.AbortWithStatusJSON(
			http.StatusRequestEntityTooLarge,
			gin.H{"message": "Filesize is too large. Max allowed filesize 1 GB."},
		)
		return
	case exceptions.UnknownFileFormat:
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "Restricted file format. Allowed file formats are:" +
					constants.GetBookFileFormatsStr(false) + "."},
		)
		return
	}

	path, err := CreateFilePath(fileHeader, constants.Books, fmt.Sprint(bookMeta.ID))
	switch err.(type) {
	case exceptions.FileAlreadyExists:
		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "File with such name already exists."},
		)
		return
	}

	var bookFile models.BookFile

	fileFormat, _ := GetFileBookFormat(fileHeader)
	bookFile.BookMetaID = bookMeta.ID
	bookFile.Filename = fileHeader.Filename
	bookFile.FileFormat = fileFormat

	if result := h.DB.Create(&bookFile); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.SaveUploadedFile(fileHeader, path)
	ctx.Status(http.StatusCreated)
}

type UpdateBookRequestBody struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
	AuthorID        uint   `json:"author_id"`
	// Genres
	// Tags
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
	// TODO: Move parse mode to config
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
