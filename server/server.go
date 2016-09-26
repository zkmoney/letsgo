package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	h := http.NewServeMux()

	h.HandleFunc("/hey", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Body: ", body)

		fmt.Fprintln(w, "hi there", string(body))
	})

	h.HandleFunc("/sup", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "sup you")
	})

	if err := http.ListenAndServe(":8024", h); err != nil {
		log.Fatal(err)
	}
}
