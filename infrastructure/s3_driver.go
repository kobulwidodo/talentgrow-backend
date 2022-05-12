package infrastructure

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type DriverS3 struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
	SessionToken    string
	BucketName      string
}

func NewS3Driver() DriverS3 {
	return DriverS3{AccessKeyId: os.Getenv("AWS_ACCESS_KEY_ID"), SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"), Region: os.Getenv("AWS_REGION"), SessionToken: os.Getenv("AWS_SESSION_TOKEN"), BucketName: os.Getenv("BUCKET_NAME")}
}

func (d *DriverS3) GetSession() (*session.Session, error) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(d.Region),
		Credentials: credentials.NewStaticCredentials(d.AccessKeyId, d.SecretAccessKey, d.SessionToken),
	})

	if err != nil {
		return sess, err
	}

	return sess, nil
}

func (d *DriverS3) UploadPublicFile(file multipart.File, fileHeader *multipart.FileHeader, userId uint) (string, error) {
	sess, err := d.GetSession()
	if err != nil {
		return "", err
	}
	nanoid, err := gonanoid.New()
	if err != nil {
		return "", err
	}
	randomKey := fmt.Sprintf("cv/%d/%s-%s", userId, nanoid, fileHeader.Filename)
	uploader := s3manager.NewUploader(sess)
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(d.BucketName),
		ACL:    aws.String("public-read"),
		Key:    aws.String(randomKey),
		Body:   file,
	})
	if err != nil {
		return "", err
	}
	return up.Location, nil
}
