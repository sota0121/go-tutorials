package webapp

import (
	"fmt"
	"log"
	"net/http"
)

// Main is the entry point of the wiki application.
func Main() {
	title1st := "1st-page"
	p1 := &Page{Title: title1st, Body: []byte("This is the 1st page.")}
	p1.save()
	p2, _ := loadPage(title1st)
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
