package handlers

import (
	"booking_v2/server/config"
	elst "booking_v2/server/elastic/user"
	user2 "booking_v2/server/models/user"
	"booking_v2/server/store"
	"booking_v2/server/utils"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(w, "login", nil)
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	values := r.Form
	login := values.Get("login")
	password := values.Get("password")
	check := values.Get("check")
	if foundUser := elst.NewRequest().QueryFilters("", login).GetUser(); foundUser != nil {
		err := bcrypt.CompareHashAndPassword([]byte(foundUser.PassHash), []byte(password+config.GlobalConfig.HashSalt))
		if err == nil {
			utils.SetSession(foundUser.Login, check, w)
			http.Redirect(w, r, "/booking", 301)
			return
		}
		_, _ = w.Write([]byte("Password is wrong"))
		return
	}
	_, _ = w.Write([]byte("Login is wrong"))
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(w, "registration", nil)
}

func PostRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	goodValues := validateRegistrationForm(r.Form)
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

func validateRegistrationForm(values url.Values) (good bool) {
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

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/booking", 302)
}

func AuthMiddleware(next http.Handler) http.Handler {
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
