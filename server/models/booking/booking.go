package booking

import "time"

type Booking struct {
	Id         string    `json:"id"`
	Number     int       `json:"number"`
	Author     string    `json:"author"`
	Time       time.Time `json:"time"`
	TimeString string    `json:"time_string"`
	Message    string    `json:"message"`
}

func New(
	id string,
	number int,
	author string,
	time time.Time,
	timeString string,
	message string,
) Booking {
	return Booking{
		Id:         id,
		Number:     number,
		Author:     author,
		Time:       time,
		TimeString: timeString,
		Message:    message,
	}
}
