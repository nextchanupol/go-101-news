package app

import (
	"log"
	"net/http"

	"github.com/nextchanupol/go-101-news/pkg/models"
	"github.com/nextchanupol/go-101-news/pkg/view"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	list, err := models.NewsList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println(list)
	view.Index(w, &view.IndexData{
		List: list,
	})
}
