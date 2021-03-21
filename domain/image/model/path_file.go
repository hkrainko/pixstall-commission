package model

import (
	"pixstall-commission/domain/file/model"
)

type PathFile struct {
	Path string
	Name string
	File model.File
}