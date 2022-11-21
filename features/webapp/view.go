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
	templNamePath    = map[string]string{
		viewTemplateName: getTmplFilePath(templatesDir, viewTemplateName),
		editTemplateName: getTmplFilePath(templatesDir, editTemplateName),
	}
	templates = template.Must(
		template.ParseFiles(
			templNamePath[viewTemplateName],
			templNamePath[editTemplateName],
		),
	)
)

func getTmplFileName(tmpl string) string {
	return fmt.Sprintf("%s.html", tmpl)
}

func getTmplFilePath(templroot, templ string) string {
	return fmt.Sprintf("%s/%s", templroot, getTmplFileName(templ))
}

func renderTemplate(w http.ResponseWriter, templ string, p *Page) {
	err := templates.ExecuteTemplate(w, getTmplFileName(templ), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, viewTemplateName, p)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		// if the page does not exist, create a new page
		p = &Page{Title: title}
	}

	renderTemplate(w, editTemplateName, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
