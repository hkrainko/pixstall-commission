package image

import (
	"context"
	model2 "pixstall-commission/domain/file/model"
)

type Repo interface {
	SaveFile(ctx context.Context, file model2.File, fileType model2.FileType) (*string, error)
	SaveFiles(ctx context.Context, files []model2.File, fileType model2.FileType) ([]string, error)
}
