package handler

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/opengraph"
	"net/http"
)

type OGP struct {
	Title       string `json:"title"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "query parameters error", http.StatusBadRequest)
		return
	}

	ogp, err := opengraph.Fetch(url)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := OGP{ogp.Title, ogp.URL.Source, ogp.Description, ogp.Image[0].URL}

	res, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(res))

}
