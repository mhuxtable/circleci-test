package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s!", r.URL.Query().Get("name"))
}

func server(port int, errorChannel chan error) {
	fmt.Println("Starting server")

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	errorChannel <- err

	fmt.Println("Listening on ", listener.Addr().String())

	s := &http.Server{
		Handler: handlers.LoggingHandler(os.Stdout, http.DefaultServeMux),
	}

	http.HandleFunc("/", testHandler)
	errorChannel <- s.Serve(listener)
}

func main() {
	serverError := make(chan error)
	go server(8080, serverError)

	for err := range serverError {
		switch err {
		case nil:
			continue
		default:
			panic(err)
			break
		}
	}
}
