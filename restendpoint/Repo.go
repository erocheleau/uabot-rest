package restendpoint

import (
	"github.com/k0kubun/pp"
	"github.com/patrickjuchli/couch"
)

var srv *couch.Server
var db *couch.Database

func RepoInit() {
	cred := couch.NewCredentials("", "")
	srv = couch.NewServer("http://127.0.0.1:5984", cred)
	db = srv.Database("orgs")
	vr, err := RepoListAllOrgs()
	if err != nil {
		for _, row := range vr.Rows {
			pp.Println(row.Value)
		}
	}
}

func RepoListAllOrgs() (*couch.ViewResult, error) {
	if srv != nil && db != nil {
		if db.HasView("org", "AllOrgs") {
			return db.Query("org", "AllOrgs", nil)
		}
	}
	return nil, nil
}

func RepoFindOrgByName(name string) Org {
	return Org{}
}

func RepoFindOrgById(id int) Org {
	return Org{}
}

func RepoCreateOrg(o Org) Org {
	return Org{}
}

func RepoDestroyOrg(id int) error {
	return nil
	//return fmt.Errorf("Could not find the Org with id of %d to delete", id)
}
