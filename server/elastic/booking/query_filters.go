package booking

import "github.com/olivere/elastic"

const (
	idField = "id"
)

type request struct {
	queryFilters QueryFilters
}

type QueryFilters struct {
	id string
}

func NewRequest() *request {
	return &request{}
}

func (r *request) QueryFilters(id string) {
	r.queryFilters = QueryFilters{
		id: id,
	}
}

func (r *request) buildSearchQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if r.queryFilters.id != "" {
		query.Must(elastic.NewTermQuery(idField, r.queryFilters.id))
	}

	return query
}
