package main

import (
	"net/http"
	"github.com/nextchanupol/go-101-news/pkg/app"
)

const port = ":8080"

func main(){
	mux := http.NewServeMux()
	app.Mount(mux)
	http.ListenAndServe(port, mux)
}