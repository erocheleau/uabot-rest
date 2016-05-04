package restendpoint

import (
	"fmt"
	"html"
	"net/http"

	"github.com/erocheleau/uabot-rest/couchclient"
)

type ViewResult struct {
	ID  string      `json:"id"`
	Key interface{} `json:"key"`
}

type ViewResponse struct {
	TotalRows int          `json:"total_rows"`
	Offset    int          `json:"offset"`
	Rows      []ViewResult `json:"rows,omitempty"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	couchconfig := couchclient.CouchConfig{url: "127.0.0.1", port: 5984}
	client := couchclient.NewClient(config)

	uuid, err := client.GetUUid()
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "Uuid %q\n", html.EscapeString(uuid))
}

func OrgsIndex(w http.ResponseWriter, r *http.Request) {
	dblist := couchclient.GetDBList()
	if len(dblist) > 0 {
		for i, dbName := range dblist {
			//pp.Fprintf(w, "Database %v: %v\n", i, dbName)
			fmt.Fprintf(w, "Database %d: %q\n", i, html.EscapeString(dbName))
		}
	}
	//w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

}

func OrgShow(w http.ResponseWriter, r *http.Request) {
	/*vars := mux.Vars(r)
	orgName := vars["orgName"]
	org := RepoFindOrgByName(orgName)
	if err := json.NewEncoder(w).Encode(org); err != nil {
		panic(err)
	}*/
}

func OrgCreate(w http.ResponseWriter, r *http.Request) {
	/*var org Org

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		panic(err)
	}

	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &org); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	o := RepoCreateOrg(org)
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(o); err != nil {
		panic(err)
	}*/
}
