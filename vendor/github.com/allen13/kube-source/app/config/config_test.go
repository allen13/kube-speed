package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	expectedAddress := "1.1.1.1"
	expectedLogging := "true"

	os.Setenv("KUBE_SOURCE_ADDRESS", expectedAddress)
	os.Setenv("KUBE_SOURCE_REQUEST_LOGGING", expectedLogging)

	Load()

	actualAddress := Get("address")

	if actualAddress != expectedAddress {
		t.Errorf("Config address failed: expected %s, got %s\n", expectedAddress,  actualAddress)
	}

	actualLogging := Get("request_logging")

	if actualLogging != expectedLogging {
		t.Errorf("Config address failed: expected %s, got %s\n", expectedLogging,  actualLogging)
	}
}