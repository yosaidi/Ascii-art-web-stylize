package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	ascii "myascii/AsciiHelper"
)

// the main handler that will redirect every root to the appropriate handler
func mainHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		ascii.MainPage(w, r)
	default:
		ascii.RenderErrorPage(w, http.StatusNotFound, "PAGE NOT FOUND", "Couldn't find this page")
	}
}

func main() {
	var err error
	style := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", style)
	ascii.Temp, err = template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal(err)
	}
	// define a global handler for all roots
	http.HandleFunc("/", mainHandler)
	// start running the server
	fmt.Println("Server started on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
