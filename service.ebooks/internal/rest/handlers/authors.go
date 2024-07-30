package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/temaxuck/WUR/service.ebooks/pkg/models"
)

type CreateAuthorRequestBody struct {
	FullName     string `json:"full_name"`
	BirthDate    string `json:"birth_date"`
	DeathDate    string `json:"death_date"`
	Description  string `json:"description"`
	WikipediaURL string `json:"wikipedia_url"`
	// Image	string
}

func (h Handler) CreateAuthor(ctx *gin.Context) {
	body := CreateAuthorRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var author models.Author

	author.FullName = body.FullName
	birth_date, _ := time.Parse("02.01.2006", body.BirthDate)
	author.BirthDate = birth_date
	death_date, _ := time.Parse("02.01.2006", body.DeathDate)
	author.DeathDate = death_date
	author.Description = body.Description
	author.WikipediaURL = body.WikipediaURL

	if result := h.DB.Create(&author); result.Error != nil {
		ctx.AbortWithError(http.StatusInternalServerError, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &author)
}

type UpdateAuthorRequestBody struct {
	FullName     string `json:"full_name"`
	BirthDate    string `json:"birth_date"`
	DeathDate    string `json:"death_date"`
	Description  string `json:"description"`
	WikipediaURL string `json:"wikipedia_url"`
	// Image	string
}

func (h Handler) UpdateAuthor(ctx *gin.Context) {
	id := ctx.Param("id")
	body := UpdateAuthorRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	defaultAuthor, err := fetchDefaultAuthor(h.DB)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if defaultAuthor.ID == author.ID {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Cannot update default author entity"})
		return
	}

	author.FullName = body.FullName
	birth_date, _ := time.Parse("02.01.2006", body.BirthDate)
	author.BirthDate = birth_date
	death_date, _ := time.Parse("02.01.2006", body.DeathDate)
	author.DeathDate = death_date
	author.Description = body.Description
	author.WikipediaURL = body.WikipediaURL

	h.DB.Save(&author)

	ctx.JSON(http.StatusOK, &author)
}

func (h Handler) GetAuthors(ctx *gin.Context) {
	var authors []models.Author

	if result := h.DB.Find(&authors); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &authors)
}

func (h Handler) GetAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusOK, &author)
}

func (h Handler) DeleteAuthor(ctx *gin.Context) {
	id := ctx.Param("id")

	defaultAuthor, err := fetchDefaultAuthor(h.DB)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var author models.Author

	if result := h.DB.First(&author, id); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if defaultAuthor.ID == author.ID {
		ctx.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Cannot delete default author entity"})
		return
	}

	h.DB.Delete(&author)

	ctx.Status(http.StatusOK)
}
