package couchclient

import (
	"time"

	couchdb "github.com/rhinoman/couchdb-go"
)

type Config struct {
	address  string
	port     int
	username string
	password string
	database string
}

type Client interface{}

type client struct {
	connection *couchdb.Connection
}

func NewClient(c Config) (Client, error) {
	var timeout = time.Duration(500 * time.Millisecond)
	conn, err := couchdb.NewConnection(c.address, c.port, timeout)
	//auth := couchdb.BasicAuth{Username: "user", Password: "password"}
	if err != nil {
		panic(err)
	}

	return &client{
		connection: conn}, nil
}
