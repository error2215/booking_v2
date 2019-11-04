package handlers

import (
	elst "booking_v2/server/elastic/user"
	user2 "booking_v2/server/models/user"
	"booking_v2/server/store"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(w, "registration", nil)
}

func PostRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	goodValues := validateUserForm(r.Form)
	if goodValues {
		user := user2.Create(r.Form.Get("name"), r.Form.Get("login"), r.Form.Get("password"))
		err := elst.NewRequest().AddUserToES(user)
		if err != nil {
			log.WithField("method", "parseAddForm").Error(err)
			_, _ = w.Write([]byte("Something went wrong, user wasn't added"))
			return
		}
		http.Redirect(w, r, "/registration", 301)
	} else {
		_, _ = w.Write([]byte("Something wrong, user wasn't added"))
	}
}

func validateUserForm(values url.Values) (good bool) {
	login := values.Get("login")
	name := values.Get("name")
	password := values.Get("password")
	password2 := values.Get("password2")
	good = true
	if user := elst.NewRequest().QueryFilters("", login).GetUser(); user == nil {
		good = false
	}
	if password != password2 {
		good = false
	}
	if len(password) < 6 {
		good = false
	}
	if len(password) > 15 {
		good = false
	}
	if len(login) < 3 {
		good = false
	}
	if len(login) > 15 {
		good = false
	}
	if len(name) < 3 {
		good = false
	}
	if len(name) > 30 {
		good = false
	}
	return good
}