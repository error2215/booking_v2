package store

import (
	"html/template"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"

	"booking_v2/server/models/user"
	"booking_v2/server/session"
)

var funcMap = template.FuncMap{}
var GlobalTemplateStore *template.Template

func init() {
	GlobalTemplateStore = template.Must(template.New("main").Funcs(funcMap).ParseGlob("templates/*.html"))
}

type pageData struct {
	Data            interface{}
	User            *user.User
	ErrorMessages   []string
	SuccessMessages []string
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

func ExecuteTemplate(r *http.Request, w http.ResponseWriter, name string, data interface{}) {
	PageData := pageData{
		Data:            data,
		User:            session.GetUser(r),
		ErrorMessages:   session.GetErrorMessages(r, w),
		SuccessMessages: session.GetSuccessMessages(r, w),
	}
	err := GlobalTemplateStore.ExecuteTemplate(w, name, PageData)
	if err != nil {
		logrus.WithField("method", "ListBookingHandler").Error(err)
	}
}
