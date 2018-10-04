/*
	gasquez.net - views
	~~~~~~~~~~~~~~~~~~~
	:license: BSD, see LICENSE for more details.
*/

package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type FooterTemplate struct {
	Year int
}

type HeaderTemplate struct {
	Title string
	PageID string
	SiteTitle string
}
type MainTemplate struct {
	Content	*Content
	Footer 	*FooterTemplate
	Header 	*HeaderTemplate
}
type MainTemplateStatic struct {
	Footer 	*FooterTemplate
	Header 	*HeaderTemplate
}
type MainMultipleTemplate struct {
	Content	*map[string]Content
	Footer 	*FooterTemplate
	Header 	*HeaderTemplate
}



var all = getContent()
var templates = template.Must(template.ParseFiles(
	"html/header.html",
	"html/footer.html",
	"html/404.html",
	"html/contact.html",
	"html/page.html",
	"html/projects.html",
))

func refreshTemplatesAndData() {
	if debug == true {
		all = getContent()
		templates = template.Must(template.ParseFiles(
			"html/header.html",
			"html/footer.html",
			"html/404.html",
			"html/contact.html",
			"html/page.html",
			"html/projects.html",
		))
	}
}

// Dead Link
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)

	data := MainTemplateStatic {
		Footer: &FooterTemplate{Year: time.Now().Year()},
		Header: &HeaderTemplate{Title: "404: Dead Link", PageID: "404", SiteTitle: siteTitle},
	}

	templates.ExecuteTemplate(w, "404.html", data)
}

// Check if we get a page, if not send NotFound
func Page(w http.ResponseWriter, r *http.Request) {
	refreshTemplatesAndData()

	vars := mux.Vars(r)


	id := "home"
	if value, ok := vars["page"]; ok {
		id = value
	}

	if post, ok := all.Pages[id]; ok {

		data := MainTemplate {
			Footer: &FooterTemplate{Year: time.Now().Year()},
			Header: &HeaderTemplate{Title: "Home", PageID: post.Slug, SiteTitle: siteTitle},
			Content: &post,
		}


		templates.ExecuteTemplate(w, "page.html", data)
	} else {
		NotFound(w, r)
	}
}

func Projects(w http.ResponseWriter, r *http.Request) {
	refreshTemplatesAndData()

	data := MainMultipleTemplate {
		Footer: &FooterTemplate{Year: time.Now().Year()},
		Header: &HeaderTemplate{Title: "My projects", PageID: "projects", SiteTitle: siteTitle},
		Content: &all.Projects,
	}

	templates.ExecuteTemplate(w, "projects.html", data)
}

// $$
func Contact(w http.ResponseWriter, r *http.Request) {
	refreshTemplatesAndData()

	data := MainTemplateStatic {
		Footer: &FooterTemplate{Year: time.Now().Year()},
		Header: &HeaderTemplate{Title: "Contact", PageID: "contact", SiteTitle: siteTitle},
	}

	templates.ExecuteTemplate(w, "contact.html", data)
}