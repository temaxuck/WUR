// TODO: Use repository pattern to implement
// different file storage systems (like local storage, S3, etc)

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
	filename := fileHeader.Filename
	extension := filepath.Ext(filename)
	if len(extension) > 0 {
		extension = extension[1:]
	}
	bookFileFormat, ok := constants.ParseBookFileFormat(extension)
	if !ok {
		return bookFileFormat, exceptions.UnknownFileFormat{
			FileName:      filename,
			FileExtension: extension,
		}
	}

	return bookFileFormat, nil
}

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

	// TODO: accept either allowed file extensions parameter
	// or enum upload type and validate file extension based
	// on the parameter value
	_, err := GetFileBookFormat(fileHeader)
	if err != nil {
		return err
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
