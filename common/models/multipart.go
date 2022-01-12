package models

import "io"

type MultipartItem struct {
	R        io.ReadCloser
	Filename string
}
