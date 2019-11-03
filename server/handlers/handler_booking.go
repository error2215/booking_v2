package handlers

import (
	"net/http"
	"strings"
	"time"

	elst "booking_v2/server/elastic/booking"
	"booking_v2/server/models/booking"
	"booking_v2/server/store"

	log "github.com/sirupsen/logrus"
)

func ListBookingHandler(w http.ResponseWriter, r *http.Request) {

}

func AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(w, "add", nil)
}

func PostAddBookingHandler(w http.ResponseWriter, r *http.Request) {
	result := parseAddForm(r)
	req := elst.Request{}
	err := req.AddBooking(result)
	if err == nil {
		_, _ = http.Get("/booking")
	} else {
		log.Error(err)
		_, _ = w.Write([]byte("Something went wrong, your booking wasn't added"))
	}

}

func parseAddForm(r *http.Request) booking.Booking {
	_ = r.ParseForm()
	timeValue := r.Form["time"][0]
	trueTime := timeValue[:len(timeValue)-11]
	trueDate := strings.ReplaceAll(timeValue[6:], "/", ".")
	neededTime, err := time.Parse("01.02.2006 15:04", trueDate+" "+trueTime)
	if err != nil {
		log.Error(err)
	}
	return booking.Booking{
		Author:  r.Form["author"][0],
		Message: r.Form["message"][0],
		Id:      elst.GetLastId() + 1,
		Time:    neededTime.Add(time.Hour * 3),
	}
}
