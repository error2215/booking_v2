package server

import (
	"net/http"

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
		r.Get("/", handlers.ListBooking) // GET /articles
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
	log.Info("Application started on port: " + config.GlobalConfig.AppPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.AppPort, r)
	if err != nil {
		log.Info(err)
	}
}
