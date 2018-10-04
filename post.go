package main

import (
	"io/ioutil"
	"strings"
	"path/filepath"
	"html/template"

	"gopkg.in/russross/blackfriday.v2" // Markdown
)

type Post struct {
	Slug 	 string // Line 0
	Type	 string // Line 1
	Title    string // Line 2
	Date     string // Line 3
	File     string 
	Body     template.HTML // Line 10+
}

func getPosts() map[string]Post {
	posts := make(map[string]Post)
	files, _ := filepath.Glob("./data/posts/*.md")

	for _, f := range files {
		fileName := strings.Replace(f, "data\\posts\\", "", -1)
		fileName = strings.Replace(fileName, ".md", "", -1)
		fileRead, _ := ioutil.ReadFile(f)

		linesStr := strings.Replace(string(fileRead), "\r", "", -1)
		lines := strings.Split(linesStr, "\n")
	
		body := strings.Join(lines[9:len(lines)], "\n")
		htmlBody := template.HTML(blackfriday.Run([]byte(body)))

		slug := string(lines[0])

		posts[slug] = Post{
			slug,
			string(lines[1]),
			string(lines[2]),
			string(lines[3]),
			fileName,
			htmlBody,
		}
	}

	return posts
}