package setup

import "github.com/srisudarshanrg/go-setup-template/server/config"

var AppConfig HandlerRepository

// HandlerRepository is the repository for the setup package which contains the app config, database and session
type HandlerRepository struct {
	AppConfig config.AppConfig
}

// RepositoryAccessSetup provides the setup package with access to the repository
func RepositoryAccessSetup(configSet HandlerRepository) {
	AppConfig = configSet
}
