package config

import "github.com/spf13/viper"

// Application configurations required to create the web application
type ApplicationConfiguration struct {
	ApplicationPort int
	EnableCORS      bool
}

var ApplicationConfig ApplicationConfiguration

// Gets config file values and sets config related to application.
// In case of no config file, default values are set
func GetAppConfig() (err error) {
	// Set configuration file
	viper.SetConfigFile("config.yml")
	// Set automatically env
	viper.AutomaticEnv()
	// read from the loaded configuration file
	err = viper.ReadInConfig()
	// Case: Error occured while reading configuration file
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// set default values in case config file is not present
			setDefault()
		} else {
			// Config file was found but another error was produced
			panic(err)
		}
	}
	// Case: Configuration read successfully
	ApplicationConfig.EnableCORS = viper.GetBool("ENABLE_CORS")
	return err
}

// Sets default values of application config
func setDefault() (err error) {
	viper.SetDefault("ENABLE_CORS", true)
	return err
}
