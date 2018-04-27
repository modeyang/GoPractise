package web

import (
	"net/http"
	"github.com/fvbock/endless"
	"fmt"
	"log"
)

func hi(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nihao!"))
}

func NewEndlessServer(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("WORLD!"))
	})

	mux.Handle("/hi", http.HandlerFunc(hi))

	err := endless.ListenAndServe(fmt.Sprintf("localhost:%d", port), mux)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server stop")
}