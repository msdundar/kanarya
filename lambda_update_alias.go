package kanarya

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// LambdaUpdateAliasResponse is used to represent the response returned from
// UpdateAlias. The main use case of this struct is to track traffic shifted
// from one version to another.
type LambdaUpdateAliasResponse struct {
	AliasArn       string
	AliasName      string
	CurrentVersion string
	NewVersion     string
	CurrentWeight  float64
}

// UpdateAlias is used to shift traffic from one version to another. version
// argument is the version to shift some traffic, and traffic argument
// stands for the amount of traffic to be shifted. For example, 0.2 means 20%
// traffic shift to the specified version.
func UpdateAlias(
	client *lambda.Lambda,
	lambdaPackage LambdaPackage,
	version string,
	traffic float64,
) (LambdaUpdateAliasResponse, error) {
	resp := LambdaUpdateAliasResponse{}

	result, err := client.UpdateAlias(
		&lambda.UpdateAliasInput{
			FunctionName: aws.String(lambdaPackage.Function.Name),
			Name:         aws.String(lambdaPackage.Alias.Name),
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
