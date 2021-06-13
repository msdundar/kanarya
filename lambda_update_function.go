package kanarya

import (
	"log"

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
	lambdaPackage LambdaPackage,
) (LambdaUpdateFunctionResponse, error) {
	resp := LambdaUpdateFunctionResponse{}

	result, err := client.UpdateFunctionCode(
		&lambda.UpdateFunctionCodeInput{
			FunctionName: aws.String(lambdaPackage.Function.Name),
			S3Bucket:     aws.String(lambdaPackage.Bucket.Name),
			S3Key:        aws.String(lambdaPackage.Bucket.Key),
		},
	)

	if err != nil {
		return resp, err
	}

	log.Println("Lambda function code updated...")

	return LambdaUpdateFunctionResponse{
		FunctionArn:      *result.FunctionArn,
		FunctionName:     *result.FunctionName,
		LastUpdateStatus: *result.LastUpdateStatus,
	}, nil
}
