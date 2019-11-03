package server

import (
	"booking_v2/server/store"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	log "github.com/sirupsen/logrus"

	"booking_v2/server/config"
	"booking_v2/server/handlers"
)

func Start() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/booking", func(r chi.Router) {
		r.Get("/", handlers.ListBookingHandler) // GET all bookings

		r.Post("/delete", handlers.DeleteBookingHandler) // POST /delete -> delete booking by id parametr

		r.Get("/add", handlers.AddBookingHandler)      // GET /add
		r.Post("/add", handlers.PostAddBookingHandler) // POST /add

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
	store.FileServer(r, "/static", http.Dir(filesDir))

	log.Info("Application started on port: " + config.GlobalConfig.AppPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.AppPort, r)
	if err != nil {
		log.Info(err)
	}
}
