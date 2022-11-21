package webapp

import (
	"fmt"
	"net/http"
	"regexp"
	"text/template"
)

var (
	// Templates
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

	// Validations
	validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
)

func getTmplFileName(tmpl string) string {
	return fmt.Sprintf("%s.html", tmpl)
}

func getTmplFilePath(templroot, templ string) string {
	return fmt.Sprintf("%s/%s", templroot, getTmplFileName(templ))
}

func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if m == nil {
		http.NotFound(w, r)
		return "", fmt.Errorf("[invalid url] cannot get title from path: %s", r.URL.Path)
	}
	// fmt.Println(m[0]) // debug --> /view/test
	// fmt.Println(m[1]) // debug --> view
	// fmt.Println(m[2]) // debug --> test
	return m[2], nil // The title is the second subexpression.
}

func renderTemplate(w http.ResponseWriter, templ string, p *Page) {
	err := templates.ExecuteTemplate(w, getTmplFileName(templ), p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// makeHandler is a closure that wraps a function with a common pattern.
// closure underlies the decorator pattern (well known in Python).
// this closure function gives the common validations and error handling functionality to the handler functions.
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// common validations and error handling for all handler functions
		title, err := getTitle(w, r)
		if err != nil {
			fmt.Println(err)
			return
		}
		// call the handler function
		fn(w, r, title)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, viewTemplateName, p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		// if the page does not exist, create a new page
		p = &Page{Title: title}
	}

	renderTemplate(w, editTemplateName, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
