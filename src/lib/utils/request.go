package utils

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

type FetchingParams struct {
	Query        bson.M
	PagingParams repository.PagingParams
}

func GetFetchingParams(w http.ResponseWriter, r *http.Request) *FetchingParams {
	params := r.URL.Query()

	// Parse query or use an empty one.
	query, err := mng.ParseQuery(params.Get(queryParamName))
	if err != nil {
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}

	// Parse 'limit' and 'skip' parameters.
	sLimit := params.Get(limitParamName)
	if sLimit == "" {
		sLimit = "0"
	}
	limit, err := strconv.Atoi(sLimit)
	if err != nil {
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}

	sSkip := params.Get(skipParamName)
	if sSkip == "" {
		sSkip = "0"
	}
	skip, err := strconv.Atoi(sSkip)
	if err != nil {
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}

	// Parse sort or use an empty one.
	sort, err := mng.ParseSort(params.Get(sortParamName))
	if err != nil {
		RenderStatusAndAbort(w, http.StatusBadRequest)
	}

	return &FetchingParams{
		Query: query,
		PagingParams: repository.PagingParams{
			Limit: limit,
			Skip:  skip,
			Sort:  sort,
		},
	}
}
