package booking

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"

	"booking_v2/server/config"
	"booking_v2/server/elastic/client"
	"booking_v2/server/models/booking"
)

func (r *Request) ListBooking() []*booking.Booking {
	query := r.BuildQuery()
	hits, err := client.GetClient().Search().
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
		err = json.Unmarshal(hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "ListBooking").Error(err)
			break
		}
		res = append(res, singleRes)
	}
	return res
}

func (r *Request) AddBooking(res booking.Booking) error {
	ourTime := time.Unix(res.Time.Local().Unix()-60*60*6, 0)
	timeStr := ourTime.Format(time.RFC850)[:len(res.Time.Format(time.RFC850))-7]
	res.TimeString = timeStr

	_, err := client.GetClient().Index().
		Index(config.GlobalConfig.BookingIndex).
		BodyJson(res).
		Id(strconv.Itoa(res.Id)).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func GetLastId() int {
	hits, err := client.GetClient().Search().
		Index(config.GlobalConfig.BookingIndex).
		Query(elastic.NewBoolQuery()).
		Sort("id", false).
		Size(1).
		Do(context.Background())
	if err != nil {
		log.Error(err)
	}
	idStr := hits.Hits.Hits[0].Id
	idInt, _ := strconv.Atoi(idStr)
	return idInt
}
