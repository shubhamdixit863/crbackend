package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"net/http"
	"strings"

	"os"
)

type FileUploadService interface {
	UploadMultipleFile([]string) ([]string, error)
}

type DefaultFileUploadService struct {
	awsS3Client *minio.Client
}

func GetFileContentType(ouput *os.File) (string, error) {

	// to sniff the content type only the first
	// 512 bytes are used.

	buf := make([]byte, 512)

	_, err := ouput.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}

func (d *DefaultFileUploadService) UploadMultipleFile(files []string) ([]string, error) {
	bucketName := os.Getenv("AWS_S3_BUCKET")
	var fileUrls []string
	//location := "fra1"
	ctx := context.Background()
	for _, file := range files {

		fileExt := file[strings.LastIndex(file, ".")+1:]

		uuidWithHyphen := uuid.New()
		uuid := fmt.Sprintf("%s.%s", strings.Replace(uuidWithHyphen.String(), "-", "", -1), fileExt)

		fileData, err := os.Open(file)
		userMetaData := map[string]string{"x-amz-acl": "public-read"}

		if err != nil {
			panic(err)
		}

		// Get the file content
		contentType, err := GetFileContentType(fileData)

		// Upload the zip file with FPutObject
		_, err = d.awsS3Client.FPutObject(ctx, bucketName, uuid, file, minio.PutObjectOptions{ContentType: contentType, UserMetadata: userMetaData})
		if err != nil {
			return nil, err
		}

		fileUrls = append(fileUrls, fmt.Sprintf("https://%s.%s/%s", bucketName, os.Getenv("END_POINT"), uuid))

	}

	return fileUrls, nil

}

func NewFileUploadService(awsS3Client *minio.Client) *DefaultFileUploadService {
	return &DefaultFileUploadService{awsS3Client: awsS3Client}

}
