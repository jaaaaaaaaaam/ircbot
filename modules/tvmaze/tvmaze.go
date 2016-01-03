package tvmaze

import (
	"encoding/json"
	"fmt"
	"github.com/jaaaaaaaaaam/ircbot/modules/tvmaze/models"
	"github.com/kennygrant/sanitize"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var showAPI = "http://api.tvmaze.com/singlesearch/shows?q="

func generateString(show models.Show) string {
	var nextLink, prevLink string
	if show.Links.Next != nil {
		nextLink = ProcessEpisode(show, "next")
	} else {
		nextLink = "No Scheduled Episode"
	}
	if show.Links.Previous != nil {
		prevLink = ProcessEpisode(show, "prev")
	} else {
		prevLink = "No Previous Episode"
	}
	str := os.Getenv("SHOWSTRING")
	r := strings.NewReplacer(
		"#ID#", string(show.ID),
		"#URL#", show.URL,
		"#showname#", show.Name,
		"#type#", show.Type,
		"#genres#", strings.Join(show.Genres[:], ", "),
		"#status#", show.Status,
		"#runtime#", string(show.Runtime),
		"#premiered#", show.Premiered,
		"#schedule.time#", string(show.Schedule.Time),
		"#schedule.days#", strings.Join(show.Schedule.Days[:], ", "),
		"#rating#", fmt.Sprintf("%f", show.Rating.Average),
		"#network.name#", show.Network.Name,
		"#network.country.name#", show.Network.Country.Name,
		"#network.country.code#", show.Network.Country.Code,
		"#network.country.timezone#", show.Network.Country.Timezone,
		"#webchannel#", show.WebChannel,
		"#externals.tvrage#", string(show.Externals.Tvrage),
		"#externals.thetvdb#", string(show.Externals.Thetvdb),
		"#image.medium#", show.Image.Medium,
		"#image.original#", show.Image.Original,
		"#summary", show.Summary,
		"#updated#", string(show.Updated), // Change this to convert to a human timestamp
		"#links.self#", show.Links.Self.Href,
		"#previousEp#", prevLink,
		"#nextEp#", nextLink,
	)
	result := r.Replace(str)
	ret, err := strconv.Unquote(`"` + result + `"`)
	if err != nil {
		fmt.Println("error replacing")
		fmt.Println(err)
	}
	return sanitize.HTML(ret)
}

// ShowLookup looks for the show and then returns a message
func ShowLookup(search string) string {
	var show models.Show
	q := strings.Join([]string{showAPI, search}, "")
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) == "" {
		return "This show doesn't exist!"
	}

	err = json.Unmarshal(body, &show)
	if err != nil {
		fmt.Println(err)
	}
	returnString := generateString(show)
	return returnString
}
