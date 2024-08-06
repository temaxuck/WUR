// TODO: Use repository pattern to implement
// different file storage systems (like local storage, S3, etc)
// TODO: Think about moving these functions to utils

package handlers

import (
	"mime/multipart"
	"path/filepath"

	"github.com/temaxuck/WUR/service.ebooks/config"
	"github.com/temaxuck/WUR/service.ebooks/internal/constants"
	"github.com/temaxuck/WUR/service.ebooks/internal/exceptions"
	"github.com/temaxuck/WUR/service.ebooks/internal/utils"
)

// TODO: Implement GetFileExtension function
func GetFileBookFormat(
	fileHeader *multipart.FileHeader,
) (constants.BookFileFormat, exceptions.EbooksError) {
	extension := GetFileExtension(fileHeader)
	bookFileFormat, ok := constants.ParseBookFileFormat(extension)
	if !ok {
		return bookFileFormat, exceptions.UnknownFileFormat{
			FileExtension: extension,
		}
	}

	return bookFileFormat, nil
}

func GetFileExtension(
	fileHeader *multipart.FileHeader,
) string {
	filename := fileHeader.Filename
	extension := filepath.Ext(filename)
	if len(extension) > 0 {
		extension = extension[1:]
	}
	return extension
}

func ValidateFile(
	fileHeader *multipart.FileHeader,
	validFileFormats []string,
) exceptions.EbooksError {
	cfg, _ := config.GetConfig()
	maxFileSize := cfg.MaxFileSize

	if fileHeader.Size > int64(maxFileSize) {
		return exceptions.FileTooLargeError{
			FileName: fileHeader.Filename,
			FileSize: uint(fileHeader.Size),
		}
	}

	extension := GetFileExtension(fileHeader)

	if !utils.ContainsStr(validFileFormats, extension) {
		return exceptions.UnknownFileFormat{FileExtension: extension}
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
