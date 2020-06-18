package grequests

import "io"

type FileUpload struct {
	// Filename is the name of the file that you wish to upload.
	FileName string

	// FileContents is happy as long as you pass it a io.ReadCloser (which most file use anyways)
	FileContents io.ReadCloser

	// FieldName is forma field name
	FieldName string

	// FileMime represents which mimetime should be sent along with the file.
	FileMime string
}
