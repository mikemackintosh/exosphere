package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	flagMessage      string
	flagMatch        string
	flagPort         string
	flagResponseCode int
)

func init() {
	flag.StringVar(&flagMessage, "o", "Hello!", "Message")
	flag.StringVar(&flagMatch, "m", "/", "Route match")
	flag.StringVar(&flagPort, "p", "8080", "Port")
	flag.IntVar(&flagResponseCode, "c", 200, "Response Code")
}

func main() {
	flag.Parse()

	srv := http.Server{
		Addr: ":" + flagPort,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(flagResponseCode)
			w.Write([]byte(flagMessage))
		}),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
