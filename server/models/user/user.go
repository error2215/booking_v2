package user

import (
	"booking_v2/server/config"
	"time"

	elst "booking_v2/server/elastic/user"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Login    string    `json:"login"`
	PassHash string    `json:"pass_hash"`
	Created  time.Time `json:"created"`
	Role     int       `json:"role"`
}

func Create(
	name string,
	login string,
	password string,
) *User {
	res := &User{}
	res.Name = name
	res.Login = login
	res.Id = elst.GetLastUserId()
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logrus.WithField("method", "models.User.New").Error(err)
		return &User{}
	}
	salt, err := bcrypt.GenerateFromPassword([]byte(config.GlobalConfig.HashSalt), 14)
	if err != nil {
		logrus.WithField("method", "models.User.New").Error(err)
		return &User{}
	}
	res.PassHash = string(pass) + string(salt)
	res.Created = time.Now()
	return res
}
