package model

import "mime/multipart"

type PathFile struct {
	Path string
	Name string
	File multipart.File
}