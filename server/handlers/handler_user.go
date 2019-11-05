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
			session.SetUser(foundUser, check, w)
			http.Redirect(w, r, "/booking", 301)
			return
		}
		session.SetErrorMessages(w, []string{"Password is wrong"})
		http.Redirect(w, r, "/login", 301)
		return
	}
	session.SetErrorMessages(w, []string{"Login is wrong"})
	http.Redirect(w, r, "/login", 301)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	store.ExecuteTemplate(r, w, "registration", nil)
}

func PostRegistrationHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	errs := validateRegistrationForm(r.Form)
	if len(errs) == 0 {
		user := model.Create(r.Form.Get("name"), r.Form.Get("login"), r.Form.Get("password"), elst.GetLastUserId())
		err := elst.NewRequest().AddUserToES(user)
		if err != nil {
			log.WithField("method", "parseAddForm").Error(err)
			session.SetErrorMessages(w, []string{"Something went wrong, please try again later"})
			http.Redirect(w, r, "/registration", 301)
			return
		}
		session.SetSuccessMessages(w, []string{"You were signed in successfully. Please login now"})
		http.Redirect(w, r, "/login", 301)
	} else {
		session.SetErrorMessages(w, errs)
		http.Redirect(w, r, "/registration", 301)
	}
}

func validateRegistrationForm(values url.Values) (errors []string) {
	login := values.Get("login")
	name := values.Get("name")
	password := values.Get("password")
	password2 := values.Get("password2")
	if user := elst.NewRequest().QueryFilters("", login).GetUser(); user != nil {
		errors = append(errors, "Such login is already used")
	}
	if password != password2 {
		errors = append(errors, "Passwords do not match")
	}
	if len(password) < 6 {
		errors = append(errors, "Password length must be more than 6 symbols")
	}
	if len(password) > 15 {
		errors = append(errors, "Password length must be less than 15 symbols")
	}
	if len(login) < 3 {
		errors = append(errors, "Login length must be more than 3 symbols")
	}
	if len(login) > 15 {
		errors = append(errors, "Login length must be less than 15 symbols")
	}
	if len(name) < 1 {
		errors = append(errors, "Name must be more than 1 symbol")
	}
	if len(name) > 30 {
		errors = append(errors, "Name length must be less than 30 symbols")
	}
	return errors
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
		user := session.GetUser(r).Name
		if user == "" {
			//TODO
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
