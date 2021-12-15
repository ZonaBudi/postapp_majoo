package config_test

import (
	"path/filepath"
	"postapp/pkg/config"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConfig(t *testing.T) {
	//load config
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "../..")
	config := &config.EnvConfig{
		FileName: "config.test",
		Path:     basepath,
	}
	err := config.ReadConfig()
	assert.NoError(t, err)
}
