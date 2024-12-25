package setup

import "github.com/srisudarshanrg/go-setup-template/server/config"

var AppConfig HandlerRepository

type HandlerRepository struct {
	AppConfig config.AppConfig
}

func RepositoryAccessSetup(configSet HandlerRepository) {
	AppConfig = configSet
}
