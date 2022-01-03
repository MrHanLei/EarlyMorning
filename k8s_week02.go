package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/healthz", Healthz)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	for k, vs := range r.Header {
		for _, v := range vs {
			w.Header().Add(k, v)
		}
	}

	w.Header().Add("VERSION", os.Getenv("VERSION"))
	log.Print("ip==", r.RemoteAddr)

	io.WriteString(w, "ok")

}

func Healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
