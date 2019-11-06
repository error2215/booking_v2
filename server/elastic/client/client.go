package client

import (
	"context"

	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"

	"booking_v2/server/config"
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

	exists, err := client.IndexExists(config.GlobalConfig.BookingIndex).Do(context.Background())
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}
	if !exists {
		// need for sorting, if there are no this fields in start mapping, app will break down, other part will be created automatically
		mapping := `{"mappings":{"properties":{"time":{"type":"date"}, "id":{"type": "long"}}}}`
		_, err = client.CreateIndex(config.GlobalConfig.BookingIndex).BodyString(mapping).Do(context.Background())
		if err != nil {
			log.WithField("method", "elastic.client.init").Fatal(err)
		}
	}
	existsUser, err := client.IndexExists(config.GlobalConfig.UserIndex).Do(context.Background())
	if err != nil {
		log.WithField("method", "elastic.client.init").Fatal(err)
	}
	if !existsUser {
		// need for sorting, if there is no this field in start mapping, app will break down, other part will be created automatically
		mapping := `{"mappings":{"properties":{"id":{"type":"long"}}}}`
		_, err = client.CreateIndex(config.GlobalConfig.UserIndex).BodyString(mapping).Do(context.Background())
		if err != nil {
			log.WithField("method", "elastic.client.init").Fatal(err)
		}
	}

	log.Info("Connection to ES cluster finished. Address: " + config.GlobalConfig.ElasticAddress)
}

func GetClient() *elastic.Client {
	return client
}
