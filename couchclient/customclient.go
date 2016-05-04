package couchclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type CouchConfig struct {
	Server string
}

type CouchClient interface{}

type couchClient struct {
	httpClient *http.Client
	server     string
}

type ViewResult struct {
	ID    string      `json:"id"`
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ViewResult `json:"rows,omitempty"`
}

func NewCouchClient(c CouchConfig) CouchClient {
	return &couchClient{
		httpClient: http.DefaultClient,
		server:     c.Server,
	}
}

func (c *couchClient) GetAllDocs(database string) (results ViewResponse, err error) {
	resp, err := http.Get(c.server + "/" + database + "/_all_docs")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return ViewResponse{}, fmt.Errorf("Error parsin json response: %v\n", err)
	}

	return results, nil
}

func (c *couchClient) GetView(database string, designdoc string, view string, params ...string) (results ViewResponse, err error) {

	fullpath := buildURL(c.server, database, "_design", designdoc, "_view", view)

	if params != nil {
		fullpath = buildURLParams(fullpath, params)
	}

	resp, err := http.Get(fullpath)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &results); err != nil {
		return ViewResponse{}, fmt.Errorf("Error parsin json response: %v\n", err)
	}

	return results, nil
}

func buildURL(base string, segments ...string) string {
	result := base
	for _, seg := range segments {
		result = result + "/" + seg
	}
	return result
}

func buildURLParams(urlBase string, params []string) string {
	strParams := ""
	for i, param := range params {
		if i > 0 {
			strParams = strParams + "&"
		}
		strParams = strParams + param
	}
	return urlBase + "?" + strParams
}
