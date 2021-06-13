package kanarya

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Client(region string) *s3.S3 {
	path_style := false

	if os.Getenv("CI") == "true" {
		path_style = true
	}

	return s3.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
		&aws.Config{
			Region:           aws.String(region),
			Endpoint:         aws.String(os.Getenv("AWS_S3_ENDPOINT")),
			S3ForcePathStyle: aws.Bool(path_style),
		},
	)
}
