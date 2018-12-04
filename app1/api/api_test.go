package api

import (
	"fmt"
	"github.com/lpxxn/gotest/app1/config"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv(config.ENV_STR, config.ENV_DEV)
	config.LoadConfig()
	fmt.Printf("Env is %s\n", os.Getenv(config.ENV_STR))
	fmt.Printf("config info %#v\n", config.ConfigInfo)

	// do something... eg: init data
	// ...

	os.Exit(m.Run())
}
