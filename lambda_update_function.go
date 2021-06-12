package kanarya

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaUpdateFunctionResponse struct {
	FunctionArn      string
	FunctionName     string
	LastUpdateStatus string
}

func UpdateFunctionCode(
	client *lambda.Lambda,
	lambda_package LambdaPackage,
) (LambdaUpdateFunctionResponse, error) {
	resp := LambdaUpdateFunctionResponse{}

	result, err := client.UpdateFunctionCode(
		&lambda.UpdateFunctionCodeInput{
			FunctionName: aws.String(lambda_package.Function.Name),
			S3Bucket:     aws.String(lambda_package.Bucket.Name),
			S3Key:        aws.String(lambda_package.Bucket.Key),
		},
	)

	if err != nil {
		return resp, err
	}

	fmt.Println("Lambda function code updated...")

	return LambdaUpdateFunctionResponse{
		FunctionArn:      *result.FunctionArn,
		FunctionName:     *result.FunctionName,
		LastUpdateStatus: *result.LastUpdateStatus,
	}, nil
}
