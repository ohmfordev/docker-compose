package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serve index.html at the root path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Println("Starting server at port 80")
	if err := http.ListenAndServe(":8001", nil); err != nil {
		fmt.Println(err)
	}
}
