package utils

import (
	"booking_v2/server/config"

	"github.com/gorilla/sessions"
)

var CookieStore = sessions.NewCookieStore([]byte(config.GlobalConfig.SessionId))
