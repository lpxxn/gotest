package config

import (
	"os"
	"testing"
)

func TestLoadConfigDev(t *testing.T) {
	os.Setenv(ENV_STR, ENV_DEV)
	LoadConfig()
	if ConfigInfo.MysqlDb.Db != "testdb" {
		t.Error(" error config info")
	}
}

func TestLoadConfigRelease(t *testing.T) {
	os.Setenv(ENV_STR, ENV_RELEASE)
	LoadConfig()
	if ConfigInfo.MysqlDb.Db != "realdb" {
		t.Error(" error config info")
	}
}