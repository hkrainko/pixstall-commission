package model

import (
	"mime/multipart"
)

type File struct {
	File   multipart.File
	Name   string
	Type   string
	Volume int64
}

type ImageFile struct {
	File
	Size
}

type Size struct {
	Width  float64 `json:"width" bson:"width"`
	Height float64 `json:"height" bson:"height"`
	Unit   string  `json:"unit" bson:"unit"`
}
