package config

import (
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

func init() {
	_ = gotenv.Load()
	viper.AutomaticEnv()

	viper.SetDefault("LAIR_ENV", "development")
	viper.SetDefault("GIN_MODE", "debug")
	viper.SetDefault("URL", "localhost:3000")
	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_USER", "postgres")
	viper.SetDefault("DB_NAME", "ag7if_development")
	viper.SetDefault("DB_PASSWORD", "postgres")
}
