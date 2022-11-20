package webapp

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	templatesDir         = "features/webapp/templates"
	viewTemplateFileName = "view.html"
	editTemplateFileName = "edit.html"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		// if the page does not exist, create a new page
		p = &Page{Title: title}
	}

	t, err := template.ParseFiles(fmt.Sprintf("%s/%s", templatesDir, editTemplateFileName))
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
