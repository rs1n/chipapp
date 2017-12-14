package services

import (
	"net/http"
	"strconv"

	"github.com/sknv/pgup"
	"github.com/sknv/pgup/orm/repository"
	"upper.io/db.v3"
)

const (
	queryParamName  = "query"
	limitParamName  = "limit"
	offsetParamName = "offset"
	orderParamName  = "order"
)

type (
	FetchingParams struct {
		Query        db.Cond
		PagingParams repository.PagingParams
	}

	Request struct{}
)

func (_ *Request) GetFetchingParamsForRequest(
	r *http.Request,
) (*FetchingParams, error) {
	params := r.URL.Query()

	// Parse query or use an empty one.
	query, err := pgup.ParseQuery(params.Get(queryParamName))
	if err != nil {
		return nil, err
	}

	// Parse 'limit' and 'offset' parameters.
	sLimit := params.Get(limitParamName)
	if sLimit == "" {
		sLimit = "0"
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		return nil, err
	}

	sOffset := params.Get(offsetParamName)
	if sOffset == "" {
		sOffset = "0"
	}
	offset, err := strconv.Atoi(sOffset)
	if err != nil {
		return nil, err
	}

	// Parse order or use an empty one.
	order, err := pgup.ParseOrder(params.Get(orderParamName))
	if err != nil {
		return nil, err
	}

	result := &FetchingParams{
		Query: query,
		PagingParams: repository.PagingParams{
			Limit:  limit,
			Offset: offset,
			Order:  order,
		},
	}
	return result, nil
}
