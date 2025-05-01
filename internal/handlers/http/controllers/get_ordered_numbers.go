package controllers

import (
	"net/http"

	"github.com/oking02/signapse-test/internal/core/ports"
)

const GetOrderedNumbersOperationID = "get_ordered_numbers"

type GetOrderedNumbers struct {
	fetcher ports.OrderedNumbersFetcher
}

func NewGetOrderedNumbers(fetcher ports.OrderedNumbersFetcher) *GetOrderedNumbers {
	return &GetOrderedNumbers{
		fetcher: fetcher,
	}
}

func (g GetOrderedNumbers) OperationID() string {
	return GetOrderedNumbersOperationID
}

func (g GetOrderedNumbers) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	orderedNumbers, err := g.fetcher.FetchOrderedNumbers(r.Context())
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			OperationID: g.OperationID(),
			Message:     "failed to get ordered numbers",
			Errors:      err.Error(),
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, Response{
		OperationID: g.OperationID(),
		Data:        orderedNumbers,
	})
}
