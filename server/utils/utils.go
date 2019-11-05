package utils

import (
	"booking_v2/server/config"
	"net/http"

	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	[]byte(config.GlobalConfig.SessionHashKey),
	[]byte(config.GlobalConfig.SessionBlockKey))

func SetSession(userName string, check string, response http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
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

func CheckUserAuth(r *http.Request) (userName string) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
