package kanarya

import (
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestPublishNewVersion(t *testing.T) {
	lambdaClient := LambdaClient(os.Getenv("AWS_REGION"))

	resp, err := PublishNewVersion(lambdaClient, testLambdaPackage)

	if err != nil {
		t.Fatalf("PublishNewVersion failed while publishing a new version. Err %v", err)
	}

	functionArn := resp.FunctionArn
	vString := strings.Split(functionArn, ":")[7]
	lastUpdateStatus := resp.LastUpdateStatus
	state := resp.State
	version := resp.Version

	lastModifiedStr := strings.Split(resp.LastModified, ".")[0]
	lastModifiedTime, err := time.Parse("2006-01-02T15:04:05", lastModifiedStr)

	if err != nil {
		t.Fatalf("Can not parse lastModifiedTime returned by PublishNewVersion. Err %v", err)
	}

	vInt, err := strconv.Atoi(vString)

	if err != nil {
		t.Fatalf("PublishNewVersion should return a new ARN for new version, but returned %v", vInt)
	}

	if lastUpdateStatus != "Successful" {
		t.Fatalf("PublishNewVersion should return a successful status, but returned %v", lastUpdateStatus)
	}

	if state != "Active" {
		t.Fatalf("PublishNewVersion should return an active state, but returned %v", state)
	}

	if version != vString {
		t.Fatalf("PublishNewVersion should return a new version number, but returned %v", version)
	}

	if lastModifiedTime.Day() != time.Now().Day() {
		t.Fatalf("PublishNewVersion should return an up-to-date last modified, but returned %v", lastModifiedTime)
	}
}
