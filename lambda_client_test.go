package kanarya

import (
	"os"
	"reflect"
	"testing"
)

// TestLambdaClient initializes a lambda client and then checks return values
// of LambdaClient function.
func TestLambdaClient(t *testing.T) {
	client := LambdaClient(os.Getenv("AWS_REGION"))

	returnValStr := reflect.TypeOf(client).String()
	expectedReturn := "*lambda.Lambda"

	if returnValStr != expectedReturn {
		t.Fatalf("LambdaClient should return %v but returned %v", expectedReturn, returnValStr)
	}

	field := reflect.TypeOf(*client).Field(0)
	fieldType := field.Type.String()
	expectedType := "*client.Client"

	if fieldType != expectedType {
		t.Fatalf("LambdaClient should be a type of %v but returned %v", expectedType, fieldType)
	}
}
