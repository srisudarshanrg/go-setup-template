package validations

import "github.com/srisudarshanrg/go-setup-template/server/config"

var appConfig config.AppConfig

// AppConfigAccessValidations provides the validations package with access to the app config
func AppConfigAccessValidations(a config.AppConfig) {
	appConfig = a
}

var db = appConfig.Database
var session = appConfig.Session
