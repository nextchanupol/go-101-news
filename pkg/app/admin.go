package app

import (
	"log"
	"net/http"

	"github.com/nextchanupol/go-101-news/pkg/models"
	"github.com/nextchanupol/go-101-news/pkg/view"
)

func adminLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		http.Redirect(w, r, "/admin/list", http.StatusSeeOther)
		return
	}
	view.AdminLogin(w, nil)
}

func adminList(w http.ResponseWriter, r *http.Request) {
	list, err := models.NewsList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	view.AdminList(w, &view.AdminListData{
		List: list,
	})
}

func adminCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		newsModel := models.News{
			Title:  r.FormValue("title"),
			Detail: r.FormValue("detail"),
		}
		err := models.CreateNews(newsModel)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		log.Println("create news")
		http.Redirect(w, r, "/admin/create", http.StatusSeeOther)
		return
	}
	view.AdminCreate(w, nil)
}

func adminEdit(w http.ResponseWriter, r *http.Request) {
	view.AdminEdit(w, nil)
}
