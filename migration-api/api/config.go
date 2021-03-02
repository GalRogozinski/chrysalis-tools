package api

import (
	"encoding/json"
	"time"

	"github.com/spf13/viper"
)

// ReadConfig reads the config.
func ReadConfig() (*Config, error) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	c := &Config{}
	if err := viper.Unmarshal(c); err != nil {
		return nil, err
	}
	return c, nil
}

type Config struct {
	ListenAddress                  string           `json:"listenAddress"`
	MaxMilestonesToQueryForEntries int              `json:"maxMilestonesToQueryForEntries"`
	MinTokenAmountForMigration     int              `json:"minTokenAmountForMigration"`
	LegacyNode                     LegacyNodeConfig `json:"legacyNode"`
	C2Node                         C2NodeConfig     `json:"c2Node"`
}

func (c *Config) JSONString() string {
	configJson, err := json.MarshalIndent(c, "", "   ")
	if err != nil {
		panic(err)
	}
	return string(configJson)
}

type LegacyNodeConfig struct {
	URI     string        `json:"uri"`
	Timeout time.Duration `json:"timeout"`
}

type C2NodeConfig struct {
	URI     string        `json:"uri"`
	Timeout time.Duration `json:"timeout"`
}
