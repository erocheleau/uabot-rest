package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	rest "github.com/erocheleau/uabot-rest/rest-endpoint"
)

func main() {
	rest.InitLogger(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	rest.LoggerInfo.Println("Server started and listening on localhost:1337/v1/")

	router := rest.NewRouter()
	log.Fatal(http.ListenAndServe(":1337", router))
}
