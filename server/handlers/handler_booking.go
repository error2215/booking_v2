package handlers

import (
	"html/template"
	"net/http"

	"github.com/sirupsen/logrus"
)

func ListBookingHandler(w http.ResponseWriter, r *http.Request) {
	filename := []string{"templates/index.html"}
	var tmpl = template.Must(template.ParseFiles(filename...))
	err := tmpl.ExecuteTemplate(w, "index", nil)
	logrus.Info(err)
}
