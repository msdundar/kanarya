package kanarya

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

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
