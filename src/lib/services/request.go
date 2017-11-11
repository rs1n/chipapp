package services

import (
	"net/http"
	"strconv"

	"github.com/sknv/chip/mng"
	"gopkg.in/mgo.v2/bson"
)

const (
	queryParamName = "query"
	limitParamName = "limit"
	skipParamName  = "skip"
)

type (
	FetchingParams struct {
		Query bson.M
		Limit int
		Skip  int
	}

	Request struct{}
)

func (_ *Request) GetFetchingParamsForRequest(
	r *http.Request,
) (*FetchingParams, error) {
	params := r.URL.Query()

	// Decode Mongo query or use an empty one.
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

	result := &FetchingParams{
		Query: query,
		Limit: limit,
		Skip:  skip,
	}
	return result, nil
}
