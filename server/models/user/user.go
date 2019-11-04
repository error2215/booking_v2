package user

import (
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	PassHash   string    `json:"pass_hash"`
	LoginToken string    `json:"login_token"`
	Created    time.Time `json:"created"`
	Role       int       `json:"role"`
}

func Create(
	name string,
	login string,
	password string,
) *User {
	res := &User{}
	res.Name = name
	res.Login = login
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		logrus.WithField("method", "models.User.New").Error(err)
		return &User{}
	}
	res.PassHash = string(bytes)
	res.LoginToken = "out"
	res.Created = time.Now()
	return res
}
