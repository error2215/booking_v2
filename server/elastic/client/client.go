package client

import (
	"booking_v2/server/config"
	"context"

	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

var GlobalClient *elastic.Client

func init() {
	GlobalClient, err := elastic.NewClient(
		elastic.SetURL(config.GlobalConfig.ElasticAddress),
	)
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}

	_, err = GlobalClient.IndexExists(config.GlobalConfig.BookingIndex).Do(context.Background())
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}
	log.Info("Connection to ES cluster finished. Address: " + config.GlobalConfig.ElasticAddress)
}
