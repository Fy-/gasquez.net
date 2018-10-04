package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux" // Routes
)

type PageData struct {
	PageTitle string
	SiteTitle string
	PageID	  string
	PagePost  Post
}
type SiteData struct {
	PageTitle string
	SiteTitle string
	PageID 	  string
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	templates, _ := template.ParseFiles(
		"html/layout.html",
		"html/404.html",
	)
	data := SiteData {
		PageTitle: "404: Dead Link",
		SiteTitle: siteTitle,
		PageID: "404",
	}

	templates.Execute(w, data)

}

func Page(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	posts := getPosts()

	templates, _ := template.ParseFiles(
		"html/layout.html",
		"html/page.html",
	)

	id := "home"
	if value, ok := vars["page"]; ok {
		id = value
	}

	if post, ok := posts[id]; ok {
		if post.Type != "Page" {
			NotFound(w, r)
		} else {
			data := PageData {
				PageTitle: "Home",
				SiteTitle: siteTitle,
				PageID: post.Slug,
				PagePost: post,
			}

			templates.Execute(w, data)
		}
	} else {
		NotFound(w, r)
	}
}

func Contact(w http.ResponseWriter, r *http.Request) {
	templates, _ := template.ParseFiles(
		"html/layout.html",
		"html/contact.html",
	)

	data := SiteData {
		PageTitle: "Contact",
		SiteTitle: siteTitle,
		PageID: "contact",
	}

	templates.Execute(w, data)
}