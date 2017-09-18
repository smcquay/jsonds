package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

var max = flag.Int("hist", 10, "number of historical events to generate")
var p = flag.Duration("period", 1*time.Minute, "how frequently to generate new alerts")
var port = flag.Int("port", 8000, "http port")

func main() {
	flag.Parse()
	s := newServer()

	// populate 10 events up front
	s.seed(*max)

	// emit period events starting now
	go s.generate(*p)

	// initialize routes, and start http server
	http.HandleFunc("/", cors(s.root))
	http.HandleFunc("/annotations", cors(s.annotations))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		log.Fatal(err)
	}
}
