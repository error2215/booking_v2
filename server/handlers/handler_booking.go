package handlers

import (
	"booking_v2/server/store"
	"net/http"
)

func ListBookingHandler(w http.ResponseWriter, r *http.Request) {

}

func AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(w, "add", nil)
}

func PostAddBookingHandler(w http.ResponseWriter, r *http.Request) {

}
