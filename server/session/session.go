package session

import (
	"net/http"
	"strconv"

	"github.com/gorilla/securecookie"

	"booking_v2/server/config"
	"booking_v2/server/models/user"
)

var cookieHandler = securecookie.New(
	[]byte(config.GlobalConfig.SessionHashKey),
	[]byte(config.GlobalConfig.SessionBlockKey))

func SetSession(user *user.User, check string, response http.ResponseWriter) {
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

func GetUserFromSession(r *http.Request) (res *user.User) {
	if cookie, err := r.Cookie("session"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("session", cookie.Value, &cookieValue); err == nil {
			res = &user.User{
				Name:  cookieValue["name"],
				Login: cookieValue["login"],
			}
			id, err := strconv.Atoi(cookieValue["id"])
			if err == nil {
				res.Id = id
			}
		}
	}
	return res
}
