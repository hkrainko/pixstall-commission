package model

import (
	"mime/multipart"
)

type File struct {
	File        multipart.File
	Name        string
	ContentType string
	FileType    FileType
	Ext         string
	Volume      int64
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

type FileType string

const (
	FileTypeMessage              = "msg"
	FileTypeCompletion           = "completion"
	FileTypeCommissionRef        = "commission-ref"
	FileTypeCommissionProofCopy  = "commission-proof-copy"
	FileTypeArtworkHidden        = "artwork-hidden"
	FileTypeArtwork              = "artwork"
	FileTypeRoof                 = "roof"
	FileTypeOpenCommission       = "open-commission"
	FileTypeOpenCommissionHidden = "open-commission-hidden"
	FileTypeProfile              = "profile"
)
