package kanarya

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaClient initializes a new lambda client that can be used in lambda
// actions. Endpoint will be set to a localstack endpoint when running tests,
// otherwise it will use the default AWS location.
func LambdaClient(region string) *lambda.Lambda {
	return lambda.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
		&aws.Config{
			Region:   aws.String(region),
			Endpoint: aws.String(os.Getenv("AWS_LAMBDA_ENDPOINT")),
		},
	)
}
