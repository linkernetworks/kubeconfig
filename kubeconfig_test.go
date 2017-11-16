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

	context := ""
	if testing_context, defined := os.LookupEnv("TEST_K8S_CONTEXT"); defined {
		context = testing_context
	}

	config, err := LoadKubeConfig(context, "")
	assert.NoError(t, err)
	assert.NotNil(t, config)
}

func TestFindKubeConfig(t *testing.T) {
	if _, defined := os.LookupEnv("TEST_K8S"); !defined {
		t.SkipNow()
		return
	}

	config, ok := FindKubeConfig()
	assert.True(t, ok)

	_, err := os.Stat(config)
	assert.NoError(t, err)
}
