package kanarya

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// testLambdaPackage is a simple LambdaPackage instance that is used
// across the test suite.
var testLambdaPackage = LambdaPackage{
	Location: "fixtures/index.zip",
	Function: LambdaFunction{
		Name: "test-lambda",
	},
	Bucket: LambdaBucket{
		Name: "test-bucket",
		Key:  "test-lambda/1.0.0/index.zip",
	},
	Alias: LambdaAlias{
		Name: "live", // alias used by clients
	},
}

// setup runs as the first step when "go test" run. It configures the test suite
// by loading environment variables defined for the test environment, and for
// tools used for testing.
func setup() {
	log.SetOutput(ioutil.Discard)

	err := godotenv.Load("test.env")

	if err != nil {
		fmt.Println("Failed to load godotenv!")
		os.Exit(1)
	}
}

// TestMain is a Go 1.14 feature. See https://golang.org/pkg/testing/#hdr-Main
// for more details.
func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
