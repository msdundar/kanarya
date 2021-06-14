package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/msdundar/kanarya"
)

var (
	awsBucket      = os.Getenv("AWS_BUCKET")
	awsRegion      = os.Getenv("AWS_REGION")
	repositoryName = os.Getenv("REPOSITORY_NAME")
	lambdaAlias    = os.Getenv("AWS_LAMBDA_ALIAS")
	lambdaName     = os.Getenv("AWS_LAMBDA_NAME")
	lambdaVersion  = os.Getenv("AWS_LAMBDA_VERSION")
)

type healthCheckQuery struct {
	Domain string
}

func main() {
	var (
		s3Client     = kanarya.S3Client(awsRegion)
		lambdaClient = kanarya.LambdaClient(awsRegion)
	)

	lambdaPackage := kanarya.LambdaPackage{
		Location: "../dist/index.zip",
		Function: kanarya.LambdaFunction{
			Name: lambdaName,
		},
		Bucket: kanarya.LambdaBucket{
			Name: awsBucket,
			Key:  fmt.Sprintf("%v/%v/%v", repositoryName, lambdaVersion, "index.zip"),
		},
		Alias: kanarya.LambdaAlias{
			Name: lambdaAlias,
		},
	}

	s3Resp, err := kanarya.UploadToS3(s3Client, lambdaPackage)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	updateResp, err := kanarya.UpdateFunctionCode(lambdaClient, lambdaPackage)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	versionResp, err := kanarya.PublishNewVersion(lambdaClient, lambdaPackage)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	newVersion := versionResp.Version

	request := healthCheckQuery{Domain: "google.com"}
	payload, err := json.Marshal(request)

	if err != nil {
		fmt.Println("Error marshalling request", err)
		os.Exit(1)
	}

	oldVersion, err := kanarya.GradualRollOut(
		lambdaClient,
		lambdaPackage,
		newVersion,
		0.1000000,
		10,
		60,
		payload,
	)

	if err != nil {
		fmt.Println(err)
		kanarya.FullRollOut(lambdaClient, lambdaPackage, oldVersion)
		os.Exit(1)
	}

	fullRolloutResp, err := kanarya.FullRollOut(lambdaClient, lambdaPackage, newVersion)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_, _, _ = s3Resp, updateResp, fullRolloutResp
}
