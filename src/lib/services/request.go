package services

import (
	"net/http"
	"strconv"

	"github.com/globalsign/mgo/bson"
	"github.com/sknv/mng"
	"github.com/sknv/mng/odm/repository"
)

const (
	limitParamName = "limit"
	skipParamName  = "skip"
	sortParamName  = "sort"
	queryParamName = "query"
)

type (
	FetchingParams struct {
		Query        bson.M
		PagingParams repository.PagingParams
	}

	Request struct{}
)

func (_ *Request) GetFetchingParamsForRequest(
	r *http.Request,
) (*FetchingParams, error) {
	params := r.URL.Query()

	// Parse query or use an empty one.
	query, err := mng.ParseQuery(params.Get(queryParamName))
	if err != nil {
		return nil, err
	}

	// Parse 'limit' and 'skip' parameters.
	sLimit := params.Get(limitParamName)
	if sLimit == "" {
		sLimit = "0"
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		return nil, err
	}

	sSkip := params.Get(skipParamName)
	if sSkip == "" {
		sSkip = "0"
	}
	skip, err := strconv.Atoi(sSkip)
	if err != nil {
		return nil, err
	}

	// Parse sort or use an empty one.
	sort, err := mng.ParseSort(params.Get(sortParamName))
	if err != nil {
		return nil, err
	}

	result := &FetchingParams{
		Query: query,
		PagingParams: repository.PagingParams{
			Limit: limit,
			Skip:  skip,
			Sort:  sort,
		},
	}
	return result, nil
}
