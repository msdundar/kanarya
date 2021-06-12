package kanarya

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

func FullRollOut(
	client *lambda.Lambda,
	lambda_package LambdaPackage,
	version string,
) (LambdaUpdateAliasResponse, error) {
	result, err := client.UpdateAlias(
		&lambda.UpdateAliasInput{
			FunctionName:    aws.String(lambda_package.Function.Name),
			Name:            aws.String(lambda_package.Alias.Name),
			FunctionVersion: aws.String(version),
			RoutingConfig:   &lambda.AliasRoutingConfiguration{},
		},
	)

	if err != nil {
		return LambdaUpdateAliasResponse{}, err
	}

	return LambdaUpdateAliasResponse{
		AliasArn:       *result.AliasArn,
		AliasName:      *result.Name,
		CurrentVersion: *result.FunctionVersion,
		NewVersion:     version,
	}, nil
}
