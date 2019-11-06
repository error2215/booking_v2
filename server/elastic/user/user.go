package user

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"

	"booking_v2/server/config"
	"booking_v2/server/elastic/client"
	model "booking_v2/server/models/user"
)

func (r *request) GetUser() *model.User {
	query := r.buildSearchQuery()
	hits, err := client.GetClient().Search().
		Index(config.GlobalConfig.UserIndex).
		Query(query).
		Size(1).
		Do(context.Background())
	if err != nil {
		log.WithField("method", "GetUserByLogin").Error(err)
	}
	if hits.TotalHits() > 0 {
		hit := hits.Hits.Hits[0]
		singleRes := &model.User{}
		err = json.Unmarshal(hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "GetUserByLogin").Error(err)
		}
		return singleRes
	}
	return nil
}

func (r *request) AddUserToES(res *model.User) error {
	_, err := client.GetClient().Index().
		Index(config.GlobalConfig.UserIndex).
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

func GetLastUserId() int {
	hits, err := client.GetClient().Search().
		Index(config.GlobalConfig.UserIndex).
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
