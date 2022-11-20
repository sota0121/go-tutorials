package webapp

import (
	"log"
	"net/http"
)

// Main is the entry point of the wiki application.
func Main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
