package restendpoint

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home")
}

func OrgsIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(orgs); err != nil {
		panic(err)
	}
}

func OrgShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orgName := vars["orgName"]
	org := RepoFindOrgByName(orgName)
	if err := json.NewEncoder(w).Encode(org); err != nil {
		panic(err)
	}
}

func OrgCreate(w http.ResponseWriter, r *http.Request) {
	var org Org

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
	}
}
