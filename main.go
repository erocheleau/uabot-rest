package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	LoggerInfo.Println("Server started and listening on localhost:1337/v1/")

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":1337", router))
}
