package restendpoint

type Org struct {
	Id                int    `json:"id"`
	Name              string `json:"name"`
	SearchEndpoint    string `json:"searchendpoint"`
	AnalyticsEndpoint string `json:"analyticsendpoint"`
}

type Orgs []Org
