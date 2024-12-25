package validations

import "github.com/srisudarshanrg/go-setup-template/server/config"

var appConfig config.AppConfig

func AppConfigAccessValidations(a config.AppConfig) {
	appConfig = a
}
