package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var port = flag.Int("p", 3000, "Port to run the HTTP Server")

	flag.Parse()

	http.Handle("/question/", HandlerWrapper(Oracle{}))

	http.Handle("/internalServerError/", HandlerWrapper(Func(func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("any Error")
	})))

	log.Printf("Listening on port %d", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		panic(err)
	}
}
