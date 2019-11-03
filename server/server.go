package server

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"booking_v2/server/config"
	"booking_v2/server/handlers"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Get("/", handlers.ListBookingHandler) // GET all bookings
		//r.Get("/{month}-{day}-{year}", listArticlesByDate) // GET /articles/01-16-2017 TODO get by date

		//r.Post("/", CreateBooking)       // POST /articles
		//r.Get("/search", searchArticles) // GET /articles/search

		// Regexp url parameters:
		//r.Get("/{articleSlug:[a-z-]+}", getArticleBySlug) // GET /articles/home-is-toronto

		// Subrouters:
		//r.Route("/{articleID}", func(r chi.Router) {
		//	r.Use(ArticleCtx)
		//	r.Get("/", getArticle)       // GET /articles/123
		//	r.Put("/", updateArticle)    // PUT /articles/123
		//	r.Delete("/", deleteArticle) // DELETE /articles/123
		//})
	})

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	FileServer(r, "/static", http.Dir(filesDir))

	log.Info("Application started on port: " + config.GlobalConfig.AppPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.AppPort, r)
	if err != nil {
		log.Info(err)
	}
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
