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

	list := models.NewsList()
	log.Println(list)
	view.Index(w, &view.IndexData{
		List: list,
	})
}
