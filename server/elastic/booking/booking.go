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
	model "booking_v2/server/models/booking"
)

func (r *request) ListBooking() []*model.Booking {
	hits, err := client.GetClient().Search().
		Query(elastic.NewBoolQuery()).
		Sort("time", true).
		Size(500).
		Index(config.GlobalConfig.BookingIndex).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "ListBooking").Error(err)
	}

	var res []*model.Booking
	if hits.TotalHits() == 0 {
		return res
	}
	for _, hit := range hits.Hits.Hits {
		singleRes := &model.Booking{}
		err = json.Unmarshal(hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "ListBooking").Error(err)
			break
		}
		res = append(res, singleRes)
	}
	return res
}

func (r *request) AddBooking(res model.Booking) error {
	ourTime := time.Unix(res.Time.Local().Unix()-60*60*6, 0)
	timeStr := ourTime.Format(time.RFC850)[:len(res.Time.Format(time.RFC850))-7]
	res.TimeString = timeStr
	res.Id = getLastId() + 1
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

func (r *request) DeleteBooking() {
	_, err := client.GetClient().Delete().
		Index(config.GlobalConfig.BookingIndex).
		Id(r.queryFilters.id).
		Refresh("true").
		Do(context.Background())
	if err != nil {
		log.Error(err)
	}
}

func getLastId() int {
	hits, err := client.GetClient().Search().
		Index(config.GlobalConfig.BookingIndex).
		Query(elastic.NewBoolQuery()).
		Sort("id", false).
		Size(1).
		Do(context.Background())
	if err != nil {
		log.Error(err)
	}
	if hits.TotalHits() == 0 {
		return 0
	}
	idStr := hits.Hits.Hits[0].Id
	idInt, _ := strconv.Atoi(idStr)
	return idInt
}
