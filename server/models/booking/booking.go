package booking

import "time"

type Booking struct {
	Id         int       `json:"id"`
	Author     string    `json:"author"`
	Time       time.Time `json:"time"`
	TimeString string    `json:"time_string"`
	Message    string    `json:"message"`
}
