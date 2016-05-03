package restendpoint

import "fmt"

var currentId int

var orgs Orgs

func init() {
	RepoCreateOrg(Org{Name: "NTO"})
	RepoCreateOrg(Org{Name: "Besttech"})
}

func RepoFindOrgByName(name string) Org {
	for _, o := range orgs {
		if o.Name == name {
			return o
		}
	}
	return Org{}
}

func RepoFindOrgById(id int) Org {
	for _, o := range orgs {
		if o.Id == id {
			return o
		}
	}
	return Org{}
}

func RepoCreateOrg(o Org) Org {
	currentId += 1
	o.Id = currentId
	orgs = append(orgs, o)
	return o
}

func RepoDestroyOrg(id int) error {
	for i, o := range orgs {
		if o.Id == id {
			orgs = append(orgs[:i], orgs[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find the Org with id of %d to delete", id)
}
