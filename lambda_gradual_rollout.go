package kanarya

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/aws/aws-sdk-go/service/lambda"
)

// GradualRollOut is the main function for gradual rollouts. It takes a version.
// argument that refers to the new version that traffic is going to be shifted
// for. traffic stands for the amount of traffic to be shifted on each step. For
// example, 0.05 stands for 5%, and 20 steps is going to be required for a full
// rollout (100/5=20). runs argument stands for number of health checks you
// would like to run on each step. sleep is number of seconds to sleep on each
// step. And finally, payload can be used as a request body to send when running
// health checks.
func GradualRollOut(
	client *lambda.Lambda,
	lambdaPackage LambdaPackage,
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
			lambdaPackage,
			version,
			sRate,
		)

		if err != nil {
			return resp.CurrentVersion, err
		}

		log.Printf(
			"Alias updated. Old version: %v, new version %v, roll out rate %v\n",
			resp.CurrentVersion,
			version,
			sRate,
		)

		for i := 0; i < runs; i++ {
			statusCodes, err := HealthCheck(client, lambdaPackage, version, payload)

			if err != nil {
				return resp.CurrentVersion, err
			}

			for _, v := range statusCodes {
				if v != 200 {
					return resp.CurrentVersion, fmt.Errorf("health check failed with %v", v)
				}
			}

			log.Println("Health checks are successful...")
		}

		time.Sleep(sleep * time.Second)
	}

	return "", nil

}
