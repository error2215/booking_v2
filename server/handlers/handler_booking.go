package handlers

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	elst "booking_v2/server/elastic/booking"
	"booking_v2/server/models/booking"
	"booking_v2/server/store"

	log "github.com/sirupsen/logrus"
)

func ListBookingHandler(w http.ResponseWriter, r *http.Request) {
	data := elst.NewRequest().ListBooking()
	store.ExecuteTemplate(r, w, "index", data)
}

func AddBookingHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(r, w, "add", nil)
}

func PostAddBookingHandler(w http.ResponseWriter, r *http.Request) {
	result := parseAddForm(r)
	err := elst.NewRequest().AddBooking(result)
	if err == nil {
		http.Redirect(w, r, "/booking", 301)
	} else {
		log.Error(err)
		_, _ = w.Write([]byte("Something went wrong, your booking wasn't added"))
	}

}

func DeleteBookingHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	id := r.Form.Get("id")
	if id != "" {
		elst.NewRequest().QueryFilters(id).DeleteBooking()
	}
	http.Redirect(w, r, "/booking", 301)
}

func DeletePastRecordsMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data := elst.NewRequest().ListBooking()
		for _, book := range data {
			unixRecordTime := book.Time.Local()
			if unixRecordTime.Add(time.Hour*-6).Unix() < time.Now().Local().Unix() {
				elst.NewRequest().QueryFilters(strconv.Itoa(book.Id)).DeleteBooking()
			}
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func parseAddForm(r *http.Request) booking.Booking {
	_ = r.ParseForm()
	timeValue := r.Form["time"][0]
	trueTime := timeValue[:len(timeValue)-11]
	trueDate := strings.ReplaceAll(timeValue[6:], "/", ".")
	neededTime, err := time.Parse("01.02.2006 15:04", trueDate+" "+trueTime)
	if err != nil {
		log.WithField("method", "parseAddForm").Error(err)
	}
	return booking.Booking{
		Author:  r.Form["author"][0],
		Message: r.Form["message"][0],
		Time:    neededTime.Add(time.Hour * 4),
	}
}
