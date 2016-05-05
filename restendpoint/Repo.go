package restendpoint

import (
	"github.com/erocheleau/uabot-rest/model"
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
	return nil, nil
}

func RepoFindOrgByName(name string) model.Org {
	return model.Org{}
}

func RepoFindOrgById(id int) model.Org {
	return model.Org{}
}

func RepoCreateOrg(o model.Org) model.Org {
	return model.Org{}
}

func RepoDestroyOrg(id int) error {
	return nil
	//return fmt.Errorf("Could not find the Org with id of %d to delete", id)
}
