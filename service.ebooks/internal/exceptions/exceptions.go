package exceptions

type EbooksError interface{}

type CantOpenFileError struct {
	EbooksError
	Err error
}

type FileTooLargeError struct {
	EbooksError
	FileName string
	FileSize uint
}

type FileAlreadyExists struct {
	EbooksError
	FileName string
}

type FileDoNotExist struct {
	EbooksError
	FileName string
}

type UnknownFileFormat struct {
	EbooksError
	FileExtension string
}
