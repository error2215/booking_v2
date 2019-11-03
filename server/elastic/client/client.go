package client

import (
	"booking_v2/server/config"
	"context"

	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
)

var client *elastic.Client

func init() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(config.GlobalConfig.ElasticAddress),
	)
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}

	_, err = client.IndexExists(config.GlobalConfig.BookingIndex).Do(context.Background())
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}
	log.Info("Connection to ES cluster finished. Address: " + config.GlobalConfig.ElasticAddress)
}

func GetClient() *elastic.Client {
	return client
}
