package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseConfigYaml(t *testing.T) {
	var yamlConfig Config
	err := ParseConfigYaml(&yamlConfig, "./config.yaml")
	assert.Nil(t, err)
	assert.Equal(t, yamlConfig.Server.Concurrency, uint(4))
	assert.Equal(t, yamlConfig.Server.SessionNumPerConnection, uint(100))
}
