package functions

import "github.com/srisudarshanrg/go-setup-template/server/config"

var appConfig config.AppConfig

// AppConfigAccessFunctions provides the functions package with access to the app config
func AppConfigAccessFunctions(a config.AppConfig) {
	appConfig = a
}

var db = appConfig.Database
