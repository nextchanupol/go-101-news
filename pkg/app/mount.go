package app

import (
	"net/http"
)

// Mount mounts handlers to mux
func Mount(mux *http.ServeMux) {
	mux.HandleFunc("/", index) // list all news
	mux.Handle("/news/", http.StripPrefix("/news", http.HandlerFunc(newsView)))
	// mux.HandleFunc("/news/", newsView) // /news/:path

	adminMux := http.NewServeMux()
	adminMux.HandleFunc("/login", adminLogin)   // /admin/login
	adminMux.HandleFunc("/list", adminList)     // /admin/list
	adminMux.HandleFunc("/create", adminCreate) // /admin/edit
	adminMux.HandleFunc("/edit", adminEdit)     // /admin/create

	mux.Handle("/admin/", http.StripPrefix("/admin", onlyAdmin(adminMux)))
}

func onlyAdmin(h http.Handler) http.Handler {
	return h
}
