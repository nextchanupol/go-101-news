package view

import (
	"net/http"

	"github.com/nextchanupol/go-101-news/pkg/models"
)

// IndexData return news data
type IndexData struct {
	List []*models.News
}

// Index renders index view
func Index(w http.ResponseWriter, data *IndexData) {
	render(tpIndex, w, data)
}

// News renders news view
func News(w http.ResponseWriter, data *models.News) {
	render(tpNews, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}

// AdminList renders admin login view
func AdminList(w http.ResponseWriter, data interface{}) {
	render(tpAdminList, w, data)
}

// AdminCreate renders admin login view
func AdminCreate(w http.ResponseWriter, data interface{}) {
	render(tpAdminCreate, w, data)
}

// AdminEdit renders admin login view
func AdminEdit(w http.ResponseWriter, data interface{}) {
	render(tpAdminEdit, w, data)
}
