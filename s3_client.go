package kanarya

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func S3Client(region string) *s3.S3 {
	return s3.New(
		session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		})),
		&aws.Config{
			Region: aws.String(region),
		},
	)
}
