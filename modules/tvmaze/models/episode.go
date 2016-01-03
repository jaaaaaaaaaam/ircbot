package models

// ShowEpisode is for data from prev/next links
type ShowEpisode struct {
	ID       int           `json:"id,omitempty"`
	URL      string        `json:"url,omitempty"`
	Name     string        `json:"name,omitempty"`
	Season   int           `json:"season,omitempty"`
	Number   int           `json:"number,omitempty"`
	Airdate  string        `json:"airdate,omitempty"`
	Airtime  string        `json:"airtime,omitempty"`
	Airstamp string        `json:"airstamp,omitempty"`
	Runtime  int           `json:"runtime,omitempty"`
	Image    *EpisodeImage `json:"image,omitempty"`
	Summary  string        `json:"summary,omitempty"`
	Links    *epLinks      `json:"_links"`
}

type epLinks struct {
	Self *epLink `json:"self"`
}

type epLink struct {
	Href string `json:"href"`
}

// EpisodeImage is the shows images
type EpisodeImage struct {
	Medium   string `json:"medium,omitempty"`
	Original string `json:"original,omitempty"`
}
