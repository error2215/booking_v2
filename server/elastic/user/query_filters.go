package user

import "github.com/olivere/elastic/v7"

const (
	idField    = "id"
	loginField = "login"
)

type request struct {
	queryFilters QueryFilters
}

type QueryFilters struct {
	id    string
	login string
}

func NewRequest() *request {
	return &request{}
}

func (r *request) QueryFilters(id string, login string) *request {
	r.queryFilters = QueryFilters{}
	if id != "" {
		r.queryFilters.id = id
	}
	if login != "" {
		r.queryFilters.login = login
	}
	return r
}

func (r *request) buildSearchQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if r.queryFilters.id != "" {
		query.Must(elastic.NewTermQuery(idField, r.queryFilters.id))
	}

	if r.queryFilters.login != "" {
		query.Must(elastic.NewTermQuery(loginField, r.queryFilters.login))
	}

	return query
}
