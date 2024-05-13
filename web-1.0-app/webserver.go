package main

import (
		"fmt"
		"log"
		"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome to the root page")
}

func main() {
		http.HandleFunc("/", index)
		log.Fatal(http.ListenAndServe(":3333", nil))
}
