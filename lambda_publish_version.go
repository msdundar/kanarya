package kanarya

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaNewVersionResponse struct {
	FunctionArn      string
	LastModified     string
	LastUpdateStatus string
	State            string
	Version          string
}

func PublishNewVersion(
	client *lambda.Lambda,
	lambdaPackage LambdaPackage,
) (LambdaNewVersionResponse, error) {
	resp := LambdaNewVersionResponse{}

	result, err := client.PublishVersion(
		&lambda.PublishVersionInput{
			FunctionName: aws.String(lambdaPackage.Function.Name),
		},
	)

	if err != nil {
		return resp, err
	}

	log.Println("Lambda version published...")

	return LambdaNewVersionResponse{
		FunctionArn:      *result.FunctionArn,
		LastModified:     *result.LastModified,
		LastUpdateStatus: *result.LastUpdateStatus,
		State:            *result.State,
		Version:          *result.Version,
	}, nil
}
