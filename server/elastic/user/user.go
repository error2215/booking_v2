package user

import (
	"booking_v2/server/config"
	"booking_v2/server/elastic/client"
	"context"
	"encoding/json"
	"strconv"

	"booking_v2/server/models/user"

	log "github.com/sirupsen/logrus"
)

func (r *request) GetUser() *user.User {
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
		singleRes := &user.User{}
		err = json.Unmarshal(hit.Source, &singleRes)
		if err != nil {
			log.WithField("method", "GetUserByLogin").Error(err)
		}
		return singleRes
	}
	return &user.User{}
}

func (r *request) AddUserToES(res *user.User) error {
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
