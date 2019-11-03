package booking

import "time"

type Booking struct {
	Id         string    `json:"_id"`
	Number     int       `json:"number"`
	Author     string    `json:"author"`
	Time       time.Time `json:"time"`
	TimeString string    `json:"time_string"`
	Message    string    `json:"message"`
}
