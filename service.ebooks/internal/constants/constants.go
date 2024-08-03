package constants

const GIGABYTE = 1_073_741_824

type UploadType string

const (
	Books  UploadType = "books"
	Covers UploadType = "covers"
	Images UploadType = "images"
)
