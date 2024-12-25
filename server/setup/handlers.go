package setup

import (
	"log"
	"net/http"

	"github.com/srisudarshanrg/go-setup-template/server/models"
)

// Home is the handler for the home page
func (app *HandlerRepository) Home(w http.ResponseWriter, r *http.Request) {
	err := RenderTemplate(w, "home.page.tmpl", models.TemplateData{})
	if err != nil {
		log.Println(err)
	}
}
