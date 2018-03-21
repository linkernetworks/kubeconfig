package kubeconfig

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLoadDefaultKubeConfig(t *testing.T) {
	if _, defined := os.LookupEnv("TEST_K8S"); !defined {
		t.SkipNow()
		return
	}

	config, err := Load("")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestFindConfig(t *testing.T) {
	if _, defined := os.LookupEnv("TEST_K8S"); !defined {
		t.SkipNow()
		return
	}

	config, ok := FindConfig()
	assert.True(t, ok)

	_, err := os.Stat(config)
	assert.NoError(t, err)
}
