package models

type Episode struct {
	ID       int    `json:"id,omitempty"`
	URL      string `json:"url,omitempty"`
	Season   int    `json:"season,omitempty"`
	Episode  int    `json:"number,omitempty"`
	Airdate  string `json:"airdate,omitempty"`
	Airstamp string `json:"airdate,omitempty"`
	Runtime  int    `json:"runtime,omitempty"`
	Image    string `json:"image,omitempty"`
	Summary  string `json:"summary,omitempty"`
}
