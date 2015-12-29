package tvmaze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"strconv"
	"os"
	"strings"
)

var showAPI = "http://api.tvmaze.com/singlesearch/shows?q="

// Show is the main JSON object
type Show struct {
	ID         int            `json:"id,omitempty"`
	URL        string         `json:"url,omitempty"`
	Name       string         `json:"name,omitempty"`
	Type       string         `json:"type,omitempty"`
	Genres     []string       `json:"genres,omitempty"`
	Status     string         `json:"status,omitempty"`
	Runtime    int            `json:"runtime,omitempty"`
	Premiered  string         `json:"premiered,omitempty"`
	Schedule   *ShowSchedule  `json:"schedule,omitempty"`
	Rating     *ShowRating    `json:"rating,omitempty"`
	Weight     int            `json:"weight,omitempty"`
	Network    *ShowNetwork   `json:"network,omitempty"`
	WebChannel string         `json:"webChannel,omitempty"`
	Externals  *ShowExternals `json:"externals,omitempty"`
	Image      *ShowImage     `json:"image,omitempty"`
	Summary    string         `json:"summary,omitempty"`
	Updated    int            `json:"updated,omitempty"`
	Links      *ShowLinks     `json:"_links,omitempty"`
}

// ShowSchedule is the schedule part of the JSON object
type ShowSchedule struct {
	Time string   `json:"time,omitempty"`
	Days []string `json:"days,omitempty"`
}

// ShowRating is the rating part of the JSON object
type ShowRating struct {
	Average float32 `json:"average,omitempty"`
}

// ShowNetwork is the network part of the JSON object
type ShowNetwork struct {
	ID      int             `json:"id,omitempty"`
	Name    string          `json:"name,omitempty"`
	Country *NetworkCountry `json:"country,omitempty"`
}

// NetworkCountry is the network country section of the JSON object
type NetworkCountry struct {
	Name     string `json:"name,omitempty"`
	Code     string `json:"code,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

// ShowExternals are the ID's of the show on different websites
type ShowExternals struct {
	Tvrage  int `json:"tvrage,omitempty"`
	Thetvdb int `json:"thetvdb,omitempty"`
}

// ShowImage is the shows images
type ShowImage struct {
	Medium   string `json:"medium,omitempty"`
	Original string `json:"original,omitempty"`
}

// ShowLinks are the links to prev/next episodes of the show
type ShowLinks struct {
	Self     *ShowLink `json:"self,omitempty"`
	Previous *ShowLink `json:"previousepisode,omitempty"`
	Next     *ShowLink `json:"nextepisode,omitempty"`
}

// ShowLink is for an individual link
type ShowLink struct {
	Href string `json:"href,omitempty"`
}

func generateString(show Show) string {
	str := os.Getenv("SHOWSTRING")
	r := strings.NewReplacer(
		"#ID#", string(show.ID),
		"#showname#", show.Name,
		"#status#", show.Status,
		"#network.name#", show.Network.Name,
	)
	result := r.Replace(str)
	return result
}

// ShowLookup looks for the show and then returns a message
func ShowLookup(search string) string {
	var show Show
	q := strings.Join([]string{showAPI, search}, "")
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if string(body) == "" {
		// The show doesn't exist
		return "This show doesn't exist!"
	}

	//err := json.Unmarshal([]byte(string(body)), $show)
	err = json.Unmarshal([]byte(string(body)), &show)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(show.Links.Self.Href)
	returnString := generateString(show)
	return returnString
}
