package config_test

import (
	"github.com/CPSSD/MDFS/config"
	"github.com/CPSSD/MDFS/utils"
	"testing"
)

func TestConfig(t *testing.T) {
	var tests = []struct {
		filename, protocol, host, port, path string
	}{
		{utils.GetUserHome() + "/.stnode/stnode_conf.json", "tcp", "localhost", "8081", utils.GetUserHome() + "/.stnode/"},
		{utils.GetUserHome() + "/.mdservice/mdservice_conf.json", "tcp", "localhost", "1994", utils.GetUserHome() + "/.mdservice/"},
	}
	for _, c := range tests {
		got := config.ParseConfiguration(c.filename)
		if got.Path != c.path || got.Host != c.host || got.Port != c.port || got.Protocol != c.protocol {
			t.Error("Configuration variables does not contain expected values")
		}
	}
}
