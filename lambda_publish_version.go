package kanarya

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaNewVersionResponse is used to represent the response returned from
// PublishNewVersion. The main use case of this struct is to check version
// number of the latest published version.
type LambdaNewVersionResponse struct {
	FunctionArn      string
	LastModified     string
	LastUpdateStatus string
	State            string
	Version          string
}

// PublishNewVersion publishes a new lambda version and returns a
// LambdaNewVersionResponse. One of the most important fields in the response is
// Version, that is the new published version and is used in later gradual
// deployment steps.
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
