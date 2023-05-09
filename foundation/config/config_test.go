package config_test

import (
	"testing"

	"github.com/brokeyourbike/microcart/foundation/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	c, err := config.Load()
	assert.NoError(t, err)
	assert.NotNil(t, c)
}
