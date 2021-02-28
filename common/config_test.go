package common

import (
	"bytes"
	"reflect"
	"testing"
)

var testConfig = []byte(
	`{
		"listenAddress": ":8000",
		"accounts": {
			"provider1": {
				"region": "us-east-1",
				"id": "key1",
				"token": "secret1",
				"defaultVPC": "vpc-123"
			},
			"provider2": {
				"region": "us-west-1",
				"id": "key2",
				"token": "secret2"
			}
		},
		"token": "SEKRET",
		"logLevel": "info",
		"org": "test"
	}`)

var brokenConfig = []byte(`{ "foobar": { "baz": "biz" }`)

func TestReadConfig(t *testing.T) {
	expectedConfig := Config{
		ListenAddress: ":8000",
		Accounts: map[string]Account{
			"provider1": {
				Region:     "us-east-1",
				Id:         "key1",
				Token:      "secret1",
				DefaultVPC: "vpc-123",
			},
			"provider2": {
				Region: "us-west-1",
				Id:     "key2",
				Token:  "secret2",
			},
		},
		Token:    "SEKRET",
		LogLevel: "info",
		Org:      "test",
	}

	actualConfig, err := ReadConfig(bytes.NewReader(testConfig))
	if err != nil {
		t.Error("Failed to read config", err)
	}

	if !reflect.DeepEqual(actualConfig, expectedConfig) {
		t.Errorf("Expected config to be %+v\n got %+v", expectedConfig, actualConfig)
	}

	_, err = ReadConfig(bytes.NewReader(brokenConfig))
	if err == nil {
		t.Error("expected error reading config, got nil")
	}
}
