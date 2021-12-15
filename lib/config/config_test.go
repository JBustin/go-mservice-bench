package config

import (
	"testing"

	"github.com/go-mservice-bench/lib/env"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	e := env.Env{
		Data: map[string]string{},
	}

	c, err := New(e)

	assert.Equal(t, nil, err, "should not raise an error.")
	assert.Equal(t, 8081, c.ServerPort, "should set a default value")
}
