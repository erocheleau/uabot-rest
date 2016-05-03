package restendpoint

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/v1/", Index},
	Route{"OrgsIndex", "GET", "/v1/orgs", OrgsIndex},
	Route{"OrgsShow", "GET", "/v1/orgs/{orgName}", OrgShow},
	Route{"OrgCreate", "POST", "/v1/orgs", OrgCreate},
}
