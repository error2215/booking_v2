package booking

import (
	"booking_v2/server/config"
	"booking_v2/server/elastic/client"
	"booking_v2/server/models/booking"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

func (r *Request) ListBooking() []*booking.Booking {
	query := r.BuildQuery()
	hits, err := client.GlobalClient.Search().
		Query(query).
		//Sort()
		Size(500).
		Index(config.GlobalConfig.BookingIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "ListBooking").Error(err)
	}

	var res []*booking.Booking
	if hits.TotalHits() == 0 {
		return res
	}
	for _, hit := range hits.Hits.Hits {
		singleRes := &booking.Booking{}
		err = json.Unmarshal(*hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "ListBooking").Error(err)
			break
		}
		res = append(res, singleRes)
	}
	return res
}
