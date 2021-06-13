package kanarya

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// FullRollOut resets the routing config and shifts whole traffic from the
// new version to the alias. It returns a LambdaUpdateAliasResponse struct.
func FullRollOut(
	client *lambda.Lambda,
	lambdaPackage LambdaPackage,
	version string,
) (LambdaUpdateAliasResponse, error) {
	result, err := client.UpdateAlias(
		&lambda.UpdateAliasInput{
			FunctionName:    aws.String(lambdaPackage.Function.Name),
			Name:            aws.String(lambdaPackage.Alias.Name),
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
