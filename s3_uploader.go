package kanarya

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3UploadResponse struct {
	ETag      string
	VersionID string
}

func UploadToS3(client *s3.S3, lambda_package LambdaPackage) (S3UploadResponse, error) {
	resp := S3UploadResponse{}

	file, err := os.Open(lambda_package.Location)

	if err != nil {
		return resp, err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	result, err := client.PutObject(
		&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(bytes.NewReader(buffer)),
			Bucket: aws.String(lambda_package.Bucket.Name),
			Key:    aws.String(lambda_package.Bucket.Key),
		},
	)

	if err != nil {
		return resp, err
	}

	fmt.Println("Lambda deployment package uploaded to S3...")

	return S3UploadResponse{
		ETag:      *result.ETag,
		VersionID: *result.VersionId,
	}, nil
}
