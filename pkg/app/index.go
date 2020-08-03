package app

import (
	"net/http"

	"github.com/nextchanupol/go-101-news/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	view.Index(w, nil)
}
