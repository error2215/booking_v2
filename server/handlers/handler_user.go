package handlers

import (
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"

	elst "booking_v2/server/elastic/user"
	model "booking_v2/server/models/user"
	"booking_v2/server/session"
	"booking_v2/server/store"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(r, w, "login", nil)
}

func PostLoginHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	values := r.Form
	login := values.Get("login")
	password := values.Get("password")
	check := values.Get("check")
	if foundUser := elst.NewRequest().QueryFilters("", login).GetUser(); foundUser != nil {
		err := bcrypt.CompareHashAndPassword([]byte(foundUser.PassHash), []byte(password))
		if err == nil {
			session.SetSession(foundUser, check, w)
			http.Redirect(w, r, "/booking", 301)
			return
		}
		log.WithField("method", "PorstLoginHandler").Error(err)
		_, _ = w.Write([]byte("Password is wrong"))
		return
	}
	_, _ = w.Write([]byte("Login is wrong"))
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(r, w, "registration", nil)
}

func PostRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	goodValues := validateRegistrationForm(r.Form)
	if goodValues {
		user := model.Create(r.Form.Get("name"), r.Form.Get("login"), r.Form.Get("password"), elst.GetLastUserId())
		err := elst.NewRequest().AddUserToES(user)
		if err != nil {
			log.WithField("method", "parseAddForm").Error(err)
			_, _ = w.Write([]byte("Something went wrong, user wasn't added"))
			return
		}
		http.Redirect(w, r, "/login", 301)
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
		user := session.GetUserFromSession(r).Name
		if user == "" {
			//TODO
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
