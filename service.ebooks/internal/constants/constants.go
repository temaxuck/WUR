package constants

import "strings"

const GIGABYTE = 1_073_741_824

type UploadType string

const (
	Books  UploadType = "books"
	Covers UploadType = "covers"
	Images UploadType = "images"
)

type BookFileFormat string

/*
When adding a new enum value here, don't forget to add it to the
"github.com/temaxuck/WUR/service.ebooks/internal/db/postgres.go"'s
Migrate function as well as to below map[string]BookFileFormat.

TODO: Find a way to fix this inconsistency
*/
const (
	LRF    BookFileFormat = "LRF"
	LRX    BookFileFormat = "LRX"
	DJVU   BookFileFormat = "DJVU"
	EPUB   BookFileFormat = "EPUB"
	FB2    BookFileFormat = "FB2"
	PDF    BookFileFormat = "PDF"
	IBOOKS BookFileFormat = "IBOOKS"
	AZW    BookFileFormat = "AZW"
	MOBI   BookFileFormat = "MOBI"
	PDB    BookFileFormat = "PDB"
	TXT    BookFileFormat = "TXT"
	RTF    BookFileFormat = "RTF"
)

var (
	book_file_formats = map[string]BookFileFormat{
		"LRF":    LRF,
		"LRX":    LRX,
		"DJVU":   DJVU,
		"EPUB":   EPUB,
		"FB2":    FB2,
		"PDF":    PDF,
		"IBOOKS": IBOOKS,
		"AZW":    AZW,
		"MOBI":   MOBI,
		"PDB":    PDB,
		"TXT":    TXT,
		"RTF":    RTF,
	}
)

func ParseBookFileFormat(f string) (BookFileFormat, bool) {
	result, ok := book_file_formats[strings.ToUpper(f)]
	return result, ok
}

func GetBookFileFormats(shouldEncloseInQuotes bool) string {
	result := ""

	mapLen := len(book_file_formats)
	i := 0

	for k := range book_file_formats {
		if shouldEncloseInQuotes {
			result += "'" + k + "'"
		} else {
			result += k
		}
		if i != mapLen-1 {
			result += ", "
		}
		i++
	}

	return result
}
