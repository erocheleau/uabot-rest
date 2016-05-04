package couchclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Client interface {
}

type Config struct {
	url  string
	port int
}

type client struct {
	httpClient *http.Client
	url        string
	port       int
}

func NewClient(c Config) Client {
	return &client{
		url:        c.url,
		port:       c.port,
		httpClient: http.DefaultClient,
	}
}

type uuidResponse struct {
	Uuids []string `json:"uuids"`
}

func (c *client) GetUUid() (string, error) {
	var buf bytes.Buffer
	var uuids uuidResponse

	resp, err := http.Get(c.url + ":" + strconv.Itoa(c.port) + "/_uuids")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &uuids); err != nil {
		return "", fmt.Errorf("Error parsin json response: %v\n", err)
	}

	return uuids.Uuids[0], nil
}
