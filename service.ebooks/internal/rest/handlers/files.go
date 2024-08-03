package handlers

import (
	"mime/multipart"

	"github.com/temaxuck/WUR/service.ebooks/config"
	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
	"github.com/temaxuck/WUR/service.ebooks/internal/exceptions"
	"github.com/temaxuck/WUR/service.ebooks/internal/utils"
)

func ValidateFile(
	fileHeader *multipart.FileHeader,
) exceptions.EbooksError {
	cfg, _ := config.GetConfig()
	maxFileSize := cfg.MaxFileSize

	if fileHeader.Size > int64(maxFileSize) {
		return exceptions.FileTooLargeError{
			FileName: fileHeader.Filename,
			FileSize: uint(fileHeader.Size),
		}
	}
	return nil
}

func CreateFilePath(
	fileHeader *multipart.FileHeader,
	uploadType constants.UploadType,
	subFolderName string,
) (string, exceptions.EbooksError) {
	path := "uploads/"

	switch uploadType {
	case constants.Books:
		path += "books/"
	case constants.Covers:
		path += "covers/"
	case constants.Images:
		path += "images/"
	}

	path += subFolderName + "/" + fileHeader.Filename

	if exists, _ := utils.PathExists(path); exists {
		return path, exceptions.FileAlreadyExists{FileName: fileHeader.Filename}
	}

	return path, nil
}
