package couchclient

import (
	"time"

	"github.com/erocheleau/uabot-rest/model"
	"github.com/rhinoman/couchdb-go"
)

type settings struct {
	server string
	port   int
}

type ViewResult struct {
	Id  string    `json:"id"`
	Key model.Org `json:"key"`
}

type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ViewResult `json:"rows,omitempty"`
}

var set = settings{server: "127.0.0.1", port: 5984}
var timeout = time.Duration(500 * time.Millisecond)

func getConnection() *couchdb.Connection {
	conn, err := couchdb.NewConnection(set.server, set.port, timeout)
	if err != nil {
		panic(err)
	}
	return conn
}

/*func GetOrgsList() (ViewResponse, error) {
	conn := getConnection()
	db := conn.SelectDB("orgstest", nil)
}*/

func GetDBList() []string {
	conn := getConnection()

	dbList, err := conn.GetDBList()
	if err != nil {
		panic(err)
	}

	return dbList
}
