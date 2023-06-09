package config_test

import (
	"os"
	"testing"

	"github.com/brokeyourbike/microcart/foundation/config"
	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	t.Parallel()

	os.Setenv("MICROCART_HTTPSERVER_ADDR", "localhost:9000")

	c, err := config.Load()
	assert.NoError(t, err)
	assert.Equal(t, "localhost:9000", c.HTTPServer.Addr)
}
