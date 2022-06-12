package tests

import (
	"BCAuth/cmd"
	"BCAuth/configuration"
	"flag"
	"fmt"
	"log"
	"os"
	"testing"
)

var (
	BCAUTH_CONFIG = "BCAUTH_CONFIG"
)

func init() {
	fmt.Println("Init tests package")
	cmd.ConfigPath = os.Getenv(BCAUTH_CONFIG)
	if cmd.ConfigPath == "" {
		fmt.Errorf("BCAUTH_CONFIG env variable was not set")
		os.Exit(1)
	}
	err := configuration.ParseConfig()
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func TestMain(m *testing.M) {
	flag.Parse()
	os.Exit(m.Run())
}
