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
	r.Use(handlers.DeletePastRecordsMiddleware)
	r.Use(handlers.AuthMiddleware)

	r.Route("/booking", func(r chi.Router) {
		r.Get("/", handlers.ListBookingHandler) // GET all bookings

		r.Post("/delete", handlers.DeleteBookingHandler) // POST /delete -> delete booking by id parameter

		r.Get("/add", handlers.AddBookingHandler)      // GET /add
		r.Post("/add", handlers.PostAddBookingHandler) // POST /add
	})

	r.Get("/login", handlers.LoginHandler)
	r.Post("/login", handlers.PostLoginHandler)

	r.Post("/logout", handlers.LogoutHandler)

	r.Get("/registration", handlers.RegistrationHandler)
	r.Post("/registration", handlers.PostRegistrationHandler)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static")
	store.FileServer(r, "/static", http.Dir(filesDir))

	log.Info("Application started on port: " + config.GlobalConfig.AppPort)
	err := http.ListenAndServe(":"+config.GlobalConfig.AppPort, r)
	if err != nil {
		log.Info(err)
	}
}
