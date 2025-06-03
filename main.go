package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "method not supported", http.StatusNotFound)
		return
	}
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.ServeFile(w, r, "./static/form.html")
		return
	} else if r.Method == "POST" {
		if r.URL.Path != "/form" {
			http.Error(w, "404 not found", http.StatusNotFound)
			return
		}

		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Fprintf(w, "POST request successful")
		name := r.FormValue("name")
		address := r.FormValue("address")
		fmt.Fprintf(w, "Hello %s, your address is %s", name, address)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)

	fmt.Printf("Starting Server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
