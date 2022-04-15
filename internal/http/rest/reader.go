package rest

import (
	"encoding/json"
	"errors"
	rdrpkg "github.com/nestorov88/rss_reader/pkg/reader"
	"net/http"
	"rss_reader_service/pkg/reader"
)

var (
	ErrWrongRequestDataFormat = errors.New("wrong request data format")
	ErrUnableReturnData       = errors.New("unable to return data")
)

//Implements http handler func
type handler struct {
	s reader.RssReaderService
}

// NewHandler return new handler with injected RssReaderService
func NewHandler(s reader.RssReaderService) *handler {
	return &handler{s}
}

type ParseUrlResponse struct {
	Items []rdrpkg.RssItem `json:"Items"`
	Error string           `json:"Error"`
}

// ParseURLs is decoding array of strings containing URLs
// parsing the urls by using reader.RssReaderService parser
// and returning ParseUrlResponse as json
func (h *handler) ParseURLs(w http.ResponseWriter, r *http.Request) {

	var urls []string

	w.Header().Set("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&urls)

	if err != nil {

		parseUrlJSONResponse(w, nil, ErrWrongRequestDataFormat)

		return
	}

	result, err := h.s.Parse(urls)

	parseUrlJSONResponse(w, result, err)
}

//parseUrlJSONResponse returns ParseUrlResponse as JSON
func parseUrlJSONResponse(w http.ResponseWriter, items []rdrpkg.RssItem, err error) {

	respData := &ParseUrlResponse{Items: items}

	if err != nil {
		respData.Error = err.Error()
	}

	err = json.NewEncoder(w).Encode(respData)

	if err != nil {
		http.Error(w, ErrUnableReturnData.Error(), http.StatusBadRequest)
	}
}
