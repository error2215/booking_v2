package store

import (
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"
)

var funcMap = template.FuncMap{}
var GlobalTemplateStore *template.Template

func init() {
	GlobalTemplateStore = template.Must(template.New("main").Funcs(funcMap).ParseGlob("templates/*.html"))
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})
}

func ExecuteTemplate(w io.Writer, name string, data interface{}) {
	err := GlobalTemplateStore.ExecuteTemplate(w, name, data)
	if err != nil {
		logrus.WithField("method", "ListBookingHandler").Error(err)
	}
}