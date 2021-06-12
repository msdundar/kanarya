package kanarya

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaUpdateAliasResponse struct {
	AliasArn       string
	AliasName      string
	CurrentVersion string
	NewVersion     string
	CurrentWeight  float64
}

func UpdateAlias(
	client *lambda.Lambda,
	lambda_package LambdaPackage,
	version string,
	traffic float64,
) (LambdaUpdateAliasResponse, error) {
	resp := LambdaUpdateAliasResponse{}

	result, err := client.UpdateAlias(
		&lambda.UpdateAliasInput{
			FunctionName: aws.String(lambda_package.Function.Name),
			Name:         aws.String(lambda_package.Alias.Name),
			RoutingConfig: &lambda.AliasRoutingConfiguration{
				AdditionalVersionWeights: map[string]*float64{
					version: aws.Float64(traffic),
				},
			},
		},
	)

	if err != nil {
		return resp, err
	}

	return LambdaUpdateAliasResponse{
		AliasArn:       *result.AliasArn,
		AliasName:      *result.Name,
		CurrentVersion: *result.FunctionVersion,
		CurrentWeight:  *result.RoutingConfig.AdditionalVersionWeights[version],
		NewVersion:     version,
	}, nil
}
