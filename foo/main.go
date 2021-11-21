package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

func middlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareOne")
		content, _ := io.ReadAll(r.Body)
		r.Body.Close() // close now
		fmt.Println(string(content))
		fmt.Println("len:", len(content))

		r.Body = io.NopCloser(bytes.NewBuffer(content))

		next.ServeHTTP(w, r)
		log.Println("Executing middlewareOne again")
	})
}

func middlewareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing middlewareTwo")
		if r.URL.Path == "/foo" {
			return
		}

		next.ServeHTTP(w, r)
		log.Println("Executing middlewareTwo again")
	})
}

func final(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing finalHandler")
	content, _ := io.ReadAll(r.Body)
	fmt.Println(string(content))
	fmt.Println("len:", len(content))
	w.Write([]byte("OK"))
}

func main() {
	mux := http.NewServeMux()

	finalHandler := http.HandlerFunc(final)
	mux.Handle("/", middlewareOne(middlewareTwo(finalHandler)))

	log.Println("Listening on :3000...")
	err := http.ListenAndServe(":3000", mux)
	log.Fatal(err)
}
