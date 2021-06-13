package kanarya

import (
	"os"
	"reflect"
	"testing"
)

// TestS3Client initializes an S3 client and then checks return values of
// S3Client function.
func TestS3Client(t *testing.T) {
	client := S3Client(os.Getenv("AWS_REGION"))

	returnValStr := reflect.TypeOf(client).String()
	expectedReturn := "*s3.S3"

	if returnValStr != expectedReturn {
		t.Fatalf("S3Client should return %v but returned %v", expectedReturn, returnValStr)
	}

	field := reflect.TypeOf(*client).Field(0)
	fieldType := field.Type.String()
	expectedType := "*client.Client"

	if fieldType != expectedType {
		t.Fatalf("S3Client should be a type of %v but returned %v", expectedType, fieldType)
	}
}
