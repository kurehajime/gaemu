package main

import (
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

func main() {
	var (
		port    = flag.Int("p", 8081, "port")
		reverse = flag.String("a", "http://localhost:8080", "Application Server URL")
		yaml    = flag.String("y", "", "app.yaml path")
	)
	flag.Parse()
	log.Printf("Listening on port %s", strconv.Itoa(*port))

	reverseURL, err := url.Parse(*reverse)
	if err != nil {
		log.Fatal(err)
	}

	yamlHandler := NewYamlHandler(reverseURL, *yaml)

	if err := http.ListenAndServe(":"+strconv.Itoa(*port), yamlHandler); err != nil {
		log.Fatal(err)
	}
}
