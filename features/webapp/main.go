package webapp

import (
	"fmt"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// Main is the entry point of the wiki application.
func Main() {
	title1st := "1st-page"
	p1 := &Page{Title: title1st, Body: []byte("This is the 1st page.")}
	p1.save()
	p2, _ := loadPage(title1st)
	fmt.Println(string(p2.Body))
}
