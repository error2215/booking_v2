package booking

import (
	"time"

	"github.com/olivere/elastic"
)

const (
	idField = "id"
)

type request struct {
	queryFilters QueryFilters
}

type QueryFilters struct {
	id   string
	time time.Time
}

func NewRequest() *request {
	return &request{}
}

func (r *request) QueryFilters(id string) *request {
	r.queryFilters = QueryFilters{
		id: id,
	}
	return r
}

func (r *request) buildSearchQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if r.queryFilters.id != "" {
		query.Must(elastic.NewTermQuery(idField, r.queryFilters.id))
	}

	return query
}
