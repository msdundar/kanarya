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
