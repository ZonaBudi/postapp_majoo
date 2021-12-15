package logger_test

import (
	"postapp/pkg/logger"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitialize(t *testing.T) {
	_, err := logger.Initialize()
	assert.NoError(t, err)
}
