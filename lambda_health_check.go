package kanarya

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/lambda"
)

type HealthCheckResponse struct {
	StatusCode int64  `json:"statusCode"`
	Body       string `json:"body"`
}

// HealthChecks runs health checks on a given lambda version or alias. It can
// take a payload argument that be used to satisfy request body expectations
// of a lambda.
func HealthCheck(
	client *lambda.Lambda,
	lambdaPackage LambdaPackage,
	version string,
	payload []byte,
) ([]int64, error) {
	var lambdaStatusCodes []int64
	var response HealthCheckResponse

	result, err := client.Invoke(&lambda.InvokeInput{
		FunctionName: aws.String(lambdaPackage.Function.Name),
		Payload:      payload,
		Qualifier:    aws.String(version),
	})

	if err != nil {
		return lambdaStatusCodes, err
	}

	err = json.Unmarshal(result.Payload, &response)

	if err != nil {
		return lambdaStatusCodes, err
	}

	lambdaStatusCodes = append(lambdaStatusCodes, *result.StatusCode, response.StatusCode)

	return lambdaStatusCodes, nil
}
