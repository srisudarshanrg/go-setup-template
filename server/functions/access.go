package functions

import "github.com/srisudarshanrg/go-setup-template/server/config"

var appConfig config.AppConfig

func AppConfigAccessFunctions(a config.AppConfig) {
	appConfig = a
}
