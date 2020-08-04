package app

import (
	"log"
	"net/http"

	"github.com/nextchanupol/go-101-news/pkg/models"
	"github.com/nextchanupol/go-101-news/pkg/view"
)

func newsView(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[1:]
	log.Println(id)
	n, err := models.GetNews(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.News(w, n)
}
