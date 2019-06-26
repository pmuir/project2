package main

import (
	"github.com/pmuir/project1/hello"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	title := "Jenkins X golang http example"

	from := ""
	if r.URL != nil {
		from = r.URL.String()
	}
	if from != "/favicon.ico" {
		log.Printf("title: %s\n", title)
	}

	hello.SayHello()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
