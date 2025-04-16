package innitialize

import (
	"fmt"

	"github.com/kyimmQ/ielts-writing-golang/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	// Init viper
	viper := viper.New()

	// Set default config file name and type
	viper.AddConfigPath("./configs/")
	viper.SetConfigName("env")
	viper.SetConfigType("yaml")

	// read config file
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read configuration %w", err))
	}

	// Load config into struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("failed to unmarshal configuration %w", err))
	}
}
