package rest

import (
	"encoding/json"
	rdrpkg "github.com/nestorov88/rss_reader/pkg/reader"
	"net/http"
	"rss_reader_service/pkg/reader"
	"time"
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
	time.Sleep(3 * time.Second)
	var urls []string

	err := json.NewDecoder(r.Body).Decode(&urls)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := h.s.Parse(urls)

	respData := &ParseUrlResponse{
		Items: result,
	}

	if err != nil {
		respData.Error = err.Error()
	}

	rawMsg, err := json.Marshal(respData)

	if err != nil {
		respData.Error = err.Error()
	}

	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

	_, err = w.Write(rawMsg)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
