package chapter1

import (
	"fmt"
	"log"
	"net/http"
)

func Ex1_11() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "URL.Path = %q\n", r.URL.Path)
}
