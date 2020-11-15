package Service

import (
	"github.com/KarthiSantha/auth/model"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func InitializeConfig(env string) model.Config {
	configuration := model.Config{}

	viper.SetConfigName("application-" + env) // config file name without extension
	viper.SetConfigType("properties")
	viper.AddConfigPath("./config/") // config file path
	viper.AutomaticEnv()             // read value ENV variable
	err := viper.ReadInConfig()

	if err != nil {
		log.WithFields(log.Fields{}).Fatal("Error in reading config file", err)
	}

	viper.SetDefault("app.port", "5000")
	// Declare var
	configuration.Port = viper.GetString("app.port")
	configuration.DBHostname = viper.GetString("db.host")
	configuration.DBPort = viper.GetString("db.port")
	configuration.DBUsername = viper.GetString("db.username")
	configuration.DBPassword = viper.GetString("db.password")
	configuration.DBDatabase = viper.GetString("db.database")
	configuration.JWTSecretKey = viper.GetString("jwt.secret_key")

	return configuration

}
