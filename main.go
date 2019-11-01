package main

import (
	"booking_v2/server"
	_ "booking_v2/server/elastic/client"
)

func main() {
	server.Start()
}
