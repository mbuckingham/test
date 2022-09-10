package main

import (
	"fmt"
	"net/http"
	"html/template"
	"time"
	"path"
)

func handle(w http.ResponseWriter, r *http.Request) {
	// You might want to move ParseGlob outside of handle so it doesn't
	// re-parse on every http request.

	
}

func main() {
	fmt.Println("http server up!")
	http.HandleFunc("/", handle)
	http.ListenAndServe(":0", nil)
}