package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

const (
	configType = "yaml"
	configName = "config"
	configPath = configName + "." + configType

	configInterval = "interval"
	configDebug    = "debug"
)

var (
	v *viper.Viper = viper.New()
)

func New() error {
	v.SetConfigType(configType)
	v.SetConfigName(configPath)
	v.AddConfigPath(".")

	v.SetDefault(configInterval, time.Second*90)
	v.SetDefault(configDebug, false)

	_ = v.ReadInConfig()
	if err := v.WriteConfigAs("config.yaml"); err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}
	return nil
}

func GetInterval() time.Duration {
	return v.GetDuration(configInterval)
}

func GetDebug() bool {
	return v.GetBool(configDebug)
}
