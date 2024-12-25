package config

import (
	"database/sql"
	"text/template"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	TemplateCache    map[string]*template.Template
	UseTemplateCache bool
	ProjectCompleted bool
	Session          *scs.SessionManager
	Database         *sql.DB
}
