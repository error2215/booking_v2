package booking

import "github.com/olivere/elastic"

const (
	idField = "id"
)

type Request struct {
	QueryFilters
}

type QueryFilters struct {
	Id string
}

func (r *Request) ParseQueryFilters() {
	r.QueryFilters = QueryFilters{}
}

func (r *Request) BuildQuery() *elastic.BoolQuery {
	query := elastic.NewBoolQuery()

	if r.QueryFilters.Id != "" {
		query.Must(elastic.NewTermQuery(idField, r.QueryFilters.Id))
	}

	return query
}
