package restendpoint

import (
	"encoding/json"
	"net/http"

	"github.com/erocheleau/uabot-rest/couchclient"
	"github.com/gorilla/mux"
)

var _couchclient couchclient.CouchClient

func Index(w http.ResponseWriter, r *http.Request) {
	config := couchclient.CouchConfig{Server: "http://localhost:5984"}
	c := couchclient.NewCouchClient(config)
	res, err := c.GetAllDocs("orgstest")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func OrgsIndex(w http.ResponseWriter, r *http.Request) {
	c := getCouchClient()

	res, err := c.GetView("orgstest", "org", "AllOrgs")
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func OrgShow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	orgName := vars["orgName"]

	var filter = "key=\"" + orgName + "\""
	c := getCouchClient()
	res, err := c.GetView("orgstest", "org", "AllOrgs", filter)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
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

func getCouchClient() couchclient.CouchClient {
	if _couchclient == nil {
		config := couchclient.CouchConfig{Server: "http://localhost:5984"}
		c := couchclient.NewCouchClient(config)
		return c
	}
	return _couchclient
}
