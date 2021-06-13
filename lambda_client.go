package kanarya

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

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
