package setup

import (
	"log"
	"net/http"
	"text/template"

	"github.com/srisudarshanrg/go-setup-template/server/models"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, templateData models.TemplateData) error {
	template, err := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	err = template.Execute(w, templateData)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
