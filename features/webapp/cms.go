package webapp

import (
	"fmt"
	"os"
)

var (
	contentsDir = "features/webapp/contents"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filepath := fmt.Sprintf("%s/%s.txt", contentsDir, p.Title)
	return os.WriteFile(filepath, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filepath := fmt.Sprintf("%s/%s.txt", contentsDir, title)
	body, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
