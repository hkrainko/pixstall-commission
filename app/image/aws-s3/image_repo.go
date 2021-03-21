package aws_s3

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	domainImage "pixstall-commission/domain/image"
	"pixstall-commission/domain/image/model"
)

type awsS3ImageRepository struct {
	s3 *s3.S3
}

const (
	BucketName = "pixstall-store-dev"
)

func NewAWSS3ImageRepository(s3 *s3.S3) domainImage.Repo {
	return &awsS3ImageRepository{
		s3: s3,
	}
}

func (a awsS3ImageRepository) SaveImage(ctx context.Context, pathImage model.PathImage) (*string, error) {
	// create buffer
	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, pathImage.Image.File.File); err != nil {
		return nil, err
	}
	uploadPath := pathImage.Path + pathImage.Name

	// convert buffer to reader
	reader := bytes.NewReader(buff.Bytes())

	// use it in `PutObjectInput`
	_, err := a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(uploadPath),
		Body:   reader,
		ContentType: aws.String("image"),
		ACL: aws.String("public-read"),  //profile should be public accessible
	})

	if err != nil {
		return nil, err
	}
	return &uploadPath, nil
}

func (a awsS3ImageRepository) SaveImages(ctx context.Context, pathImages []model.PathImage) ([]string, error) {

	var resultPaths []string
	for _, pathImage := range pathImages {
		// create buffer
		buff := new(bytes.Buffer)
		if _, err := io.Copy(buff, pathImage.Image.File.File); err != nil {
			return nil, err
		}
		uploadPath := pathImage.Path + pathImage.Name
		// convert buffer to reader
		reader := bytes.NewReader(buff.Bytes())

		// use it in `PutObjectInput`
		_, err := a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
			Bucket: aws.String(BucketName),
			Key:    aws.String(uploadPath),
			Body:   reader,
			ContentType: aws.String("image"),
			ACL: aws.String("public-read"),  //profile should be public accessible
		})
		if err == nil {
			resultPaths = append(resultPaths, uploadPath)
		}
	}
	return resultPaths, nil
}

func (a awsS3ImageRepository) SaveFile(ctx context.Context, pathFile model.PathFile) (*string, error) {
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, pathFile.File.File); err != nil {
		return nil, err
	}
	uploadPath := pathFile.Path + pathFile.Name

	// convert buffer to reader
	reader := bytes.NewReader(buf.Bytes())

	// use it in `PutObjectInput`
	_, err := a.s3.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(BucketName),
		Key:    aws.String(uploadPath),
		Body:   reader,
		//ContentType: aws.String("image"),
		ACL: aws.String("public-read"),  //profile should be public accessible
	})

	if err != nil {
		return nil, err
	}
	return &uploadPath, nil
}