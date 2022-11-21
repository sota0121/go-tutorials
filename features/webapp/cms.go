package webapp

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	contentsDir = "features/webapp/contents"
)

type Page struct {
	Title string
	Body  []byte
}

type Pages []*Page

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

func loadPages() (Pages, error) {
	var pages Pages
	filenames, err := filepath.Glob(filepath.Join(contentsDir, "*.txt"))
	if err != nil {
		return nil, err
	}
	for _, filename := range filenames {
		title := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
		p, err := loadPage(title)
		if err != nil {
			return nil, err
		}
		pages = append(pages, p)
	}
	return pages, nil
}
