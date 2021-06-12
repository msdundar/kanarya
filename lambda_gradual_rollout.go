package kanarya

import (
	"fmt"
	"math"
	"time"

	"github.com/aws/aws-sdk-go/service/lambda"
)

func GradualRollOut(
	client *lambda.Lambda,
	lambda_package LambdaPackage,
	version string,
	traffic float64,
	runs int,
	sleep time.Duration,
	payload []byte,
) (string, error) {

	for rate := traffic; rate <= 1.0; rate += traffic {
		sRate := math.Round(rate*100) / 100

		resp, err := UpdateAlias(
			client,
			lambda_package,
			version,
			sRate,
		)

		if err != nil {
			return resp.CurrentVersion, err
		}

		fmt.Printf(
			"Alias updated. Old version: %v, new version %v, roll out rate %v\n",
			resp.CurrentVersion,
			version,
			sRate,
		)

		for i := 0; i < runs; i++ {
			statusCodes, err := HealthCheck(client, lambda_package, version, payload)

			if err != nil {
				return resp.CurrentVersion, err
			}

			for _, v := range statusCodes {
				if v != 200 {
					return resp.CurrentVersion, fmt.Errorf("health check failed with %v", v)
				}
			}

			fmt.Println("Health checks are successful...")
		}

		time.Sleep(sleep * time.Second)
	}

	return "", nil

}
