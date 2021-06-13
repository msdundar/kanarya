package kanarya

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

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

func setup() {
	log.SetOutput(ioutil.Discard)

	err := godotenv.Load("test.env")

	if err != nil {
		fmt.Println("Failed to load godotenv!")
		os.Exit(1)
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
