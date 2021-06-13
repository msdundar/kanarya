package kanarya

import (
	"os"
	"reflect"
	"testing"
)

// TestS3Uploader uploads a lambda package to an S3 bucket and tests the
// response.
func TestS3Uploader(t *testing.T) {
	s3Client := S3Client(os.Getenv("AWS_REGION"))

	resp, err := UploadToS3(s3Client, testLambdaPackage)

	if err != nil {
		t.Fatalf("UploadToS3 failed while uploading the file to S3. Err %v", err)
	}

	eTag := resp.ETag

	if reflect.TypeOf(eTag).Kind() != reflect.String {
		t.Fatalf("UploadToS3 should return a string eTag, but it returned %v", eTag)
	}

	if len(eTag) != 34 {
		t.Fatalf("UploadToS3 should return a 34 char-long string eTag, but it returned %v", eTag)
	}
}
