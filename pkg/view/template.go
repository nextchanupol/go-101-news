package view

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/css"
	"github.com/tdewolff/minify/html"
	"github.com/tdewolff/minify/js"
)

// refactor parseTemplate
var (
	tpIndex      = parseTemplate("root.tmpl", "index.tmpl")
	tpAdminLogin = parseTemplate("root.tmpl", "admin/login.tmpl")
)

var m = minify.New()

const templateDir = "template"

func init() {
	m.AddFunc("text/html", html.Minify)
	m.AddFunc("text/css", css.Minify)
	m.AddFunc("text/javascript", js.Minify)
}

func joinTemplateDir(files ...string) []string {
	r := make([]string, len(files))
	for i, file := range files {
		r[i] = filepath.Join(templateDir, file)
	}
	return r
}

func parseTemplate(files ...string) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{})
	_, err := t.ParseFiles(joinTemplateDir(files...)...)
	if err != nil {
		panic(err)
	}
	t = t.Lookup("root")
	return t
}

/*

var (
	tpIndex = template.New("")
)

func init() {
	tpIndex.Funcs(template.FuncMap{})
	_, err := tpIndex.ParseFiles("template/root.tmpl", "template/index.tmpl")

	if err != nil {
		panic(err)
	}

	tpIndex = tpIndex.Lookup("root")
}*/

func render(t *template.Template, w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
		return
	}
	m.Minify("text/html", w, &buf)
}

// Index renders index view
func Index(w http.ResponseWriter, data interface{}) {
	render(tpIndex, w, data)
}

// AdminLogin renders admin login view
func AdminLogin(w http.ResponseWriter, data interface{}) {
	render(tpAdminLogin, w, data)
}
