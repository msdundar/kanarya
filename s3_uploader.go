package kanarya

import (
	"bytes"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// An S3UploadResponse is used to represent structure of a response returned
// from an UploadToS3 call. It supports S3 buckets with versioning enabled or
// disabled.
type S3UploadResponse struct {
	ETag string
}

// UploadToS3 uploads a given lambda package to an S3 bucket. Later, the object
// in the bucket can be used for deploying the lambda function.
func UploadToS3(client *s3.S3, lambdaPackage LambdaPackage) (S3UploadResponse, error) {
	resp := S3UploadResponse{}

	file, err := os.Open(lambdaPackage.Location)

	if err != nil {
		return resp, err
	}

	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	_, err = file.Read(buffer)

	if err != nil {
		return resp, err
	}

	result, err := client.PutObject(
		&s3.PutObjectInput{
			Body:   aws.ReadSeekCloser(bytes.NewReader(buffer)),
			Bucket: aws.String(lambdaPackage.Bucket.Name),
			Key:    aws.String(lambdaPackage.Bucket.Key),
		},
	)

	if err != nil {
		return resp, err
	}

	log.Println("Lambda deployment package uploaded to S3...")

	return S3UploadResponse{
		ETag: *result.ETag,
	}, nil
}
