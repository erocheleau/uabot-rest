package model

type Org struct {
	ID                string `json:"_id"`
	Rev               string `json:"_rev"`
	Name              string `json:"name"`
	SearchEndpoint    string `json:"searchendpoint"`
	AnalyticsEndpoint string `json:"analyticsendpoint"`
}

type Orgs []Org
