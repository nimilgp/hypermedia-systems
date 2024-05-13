package main

import (
		"fmt"
		"log"
		"net/http"
)

func contactHandler(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to the contact page")
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
		//fmt.Fprintf(w, "welcome to the root page")
		http.Redirect(w, r, "/contact", http.StatusMovedPermanently)
}

func main() {
		http.HandleFunc("/", indexHandler)
		http.HandleFunc("/contact", contactHandler)
		log.Fatal(http.ListenAndServe(":3333", nil))
}
