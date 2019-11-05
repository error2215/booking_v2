package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"

	"booking_v2/server/config"
	"booking_v2/server/models/user"
)

var cookieHandler = securecookie.New(
	[]byte(config.GlobalConfig.SessionHashKey),
	[]byte(config.GlobalConfig.SessionBlockKey))

func SetUser(user *user.User, check string, response http.ResponseWriter) {
	value := map[string]string{
		"login": user.Login,
		"id":    strconv.Itoa(user.Id),
		"name":  user.Name,
	}
	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		if check == "true" {
			cookie.MaxAge = 60 * 60 * 24 * 30
		}
		http.SetCookie(response, cookie)
	}
}

func GetUser(r *http.Request) *user.User {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			res := &user.User{
				Name:  cookieValue["name"],
				Login: cookieValue["login"],
			}
			id, err := strconv.Atoi(cookieValue["id"])
			if err == nil {
				res.Id = id
			}
			return res
		}
	}
	return &user.User{
		Id:       0,
		Name:     "",
		Login:    "",
		PassHash: "",
		Created:  time.Time{},
		Role:     0,
	}
}

func SetSuccessMessages(w http.ResponseWriter, msgs []string) {
	value := map[string]string{
		"successMessages": strings.Join(msgs, "|"),
	}
	if encoded, err := cookieHandler.Encode("successMessages", value); err == nil {
		cookie := &http.Cookie{
			Name:  "successMessages",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func GetSuccessMessages(r *http.Request, w http.ResponseWriter) []string {
	if cookie, err := r.Cookie("successMessages"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("successMessages", cookie.Value, &cookieValue); err == nil {
			cookie := &http.Cookie{
				Name:    "successMessages",
				Value:   "",
				Path:    "/",
				MaxAge:  -1,
				Expires: time.Unix(1, 0),
			}
			http.SetCookie(w, cookie)
			return strings.Split(cookieValue["successMessages"], "|")
		}
	}
	return []string{}
}

func SetErrorMessages(w http.ResponseWriter, msgs []string) {
	value := map[string]string{
		"errorMessages": strings.Join(msgs, "|"),
	}
	if encoded, err := cookieHandler.Encode("errorMessages", value); err == nil {
		cookie := &http.Cookie{
			Name:  "errorMessages",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func GetErrorMessages(r *http.Request, w http.ResponseWriter) []string {
	if cookie, err := r.Cookie("errorMessages"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("errorMessages", cookie.Value, &cookieValue); err == nil {
			cookie := &http.Cookie{
				Name:    "errorMessages",
				Value:   "",
				Path:    "/",
				MaxAge:  -1,
				Expires: time.Unix(1, 0),
			}
			http.SetCookie(w, cookie)
			return strings.Split(cookieValue["errorMessages"], "|")
		}
	}
	return []string{}
}
