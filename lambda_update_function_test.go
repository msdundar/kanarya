package kanarya

import (
	"os"
	"testing"
)

func TestUpdateFunctionCode(t *testing.T) {
	lambdaClient := LambdaClient(os.Getenv("AWS_REGION"))

	resp, err := UpdateFunctionCode(lambdaClient, testLambdaPackage)

	if err != nil {
		t.Fatalf("UpdateFunctionCode failed while updating the lambda function. Err %v", err)
	}

	functionArn := resp.FunctionArn
	functionName := resp.FunctionName
	lastUpdateStatus := resp.LastUpdateStatus

	if functionArn != "arn:aws:lambda:us-east-1:000000000000:function:test-lambda" {
		t.Fatalf("UpdateFunctionCode should return a proper function ARN but it returned %v", functionArn)
	}

	if functionName != "test-lambda" {
		t.Fatalf("UpdateFunctionCode should return test-lambda as function name but it returned %v", functionName)
	}

	if lastUpdateStatus != "Successful" {
		t.Fatalf("UpdateFunctionCode should return Successful as status but it returned %v", lastUpdateStatus)
	}
}
