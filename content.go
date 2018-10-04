/*
	gasquez.net - posts and pages
	~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
	:license: BSD, see LICENSE for more details.
*/

package main

import (
	"io/ioutil"
	"strings"
	"path/filepath"
	"html/template"

	"gopkg.in/russross/blackfriday.v2"
)

// Simple struct describing a post or a page.
type Content struct {
	Slug 	 string // Line 0
	Type	 string // Line 1
	Title    string // Line 2
	Date     string // Line 3
	File     string 
	Body     template.HTML // Line 10+
}

type Data struct {
	All			map[string]Content
	Pages		map[string]Content
	Projects	map[string]Content
	Posts		map[string]Content
}

// Get all posts from data/posts in a dict (and transform markdown to html)
// It's not safe but it's ok. I'm not completely stupid... I think.
func getContent() Data {
	all := make(map[string]Content)
	pages := make(map[string]Content)
	posts := make(map[string]Content)
	projects := make(map[string]Content)

	files, _ := filepath.Glob("./data/*/*.md")

	for _, f := range files {
		fileName := strings.Replace(f, "data\\posts\\", "", -1)
		fileName = strings.Replace(fileName, ".md", "", -1)
		fileRead, _ := ioutil.ReadFile(f)

		linesStr := strings.Replace(string(fileRead), "\r", "", -1)
		lines := strings.Split(linesStr, "\n")
	
		body := strings.Join(lines[9:len(lines)], "\n")
		htmlBody := template.HTML(blackfriday.Run([]byte(body)))

		slug := string(lines[0])

		all[slug] = Content{
			slug,
			string(lines[1]),
			string(lines[2]),
			string(lines[3]),
			fileName,
			htmlBody,
		}

		if all[slug].Type == "Page" {
			pages[slug] = all[slug]
		} else if all[slug].Type == "Project" {
			projects[slug] = all[slug]
		} else if all[slug].Type == "Post" {
			posts[slug] = all[slug]
		}
	}

	data := Data {
		All: all,
		Pages: pages,
		Posts: posts,
		Projects: projects,
	}

	return data
}