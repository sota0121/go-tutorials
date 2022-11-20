package webapp

import (
	"fmt"
	"net/http"
	"text/template"
)

var (
	templatesDir     = "features/webapp/templates"
	viewTemplateName = "view"
	editTemplateName = "edit"
)

func renderTemplate(w http.ResponseWriter, templroot, templ string, p *Page) {
	t, err := template.ParseFiles(fmt.Sprintf("%s/%s.html", templroot, templ))
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	t.Execute(w, p)
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	renderTemplate(w, templatesDir, viewTemplateName, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		// if the page does not exist, create a new page
		p = &Page{Title: title}
	}

	renderTemplate(w, templatesDir, editTemplateName, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
