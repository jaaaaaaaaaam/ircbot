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

func generateEpisodeString(ep models.Episode) string {
	str := os.Getenv("EPSTRING")
	fmt.Println(ep.Airdate)
	r := strings.NewReplacer(
		"#ID#", string(ep.ID),
		"#URL#", ep.URL,
		"#name#", ep.Name,
		"#season#", strconv.Itoa(ep.Season),
		"#episode#", strconv.Itoa(ep.Episode),
		"#airdate#", ep.Airdate,
		"#airtime#", ep.Airtime,
		"#airstamp#", ep.Airstamp,
		"#runtime#", string(ep.Runtime),
	)
	result := r.Replace(str)
	ret, err := strconv.Unquote(`"` + result + `"`)
	if err != nil {
		fmt.Println("error replacing")
		fmt.Println(err)
	}
	return sanitize.HTML(ret)
}

// ProcessEpisode processes either the next/prev episode data
func ProcessEpisode(data models.Show, order string) string {
	var ep models.Episode
	if order == "prev" {
		if data.Links.Previous != nil {
			resp, err := http.Get(data.Links.Previous.Href)
			if err != nil {
				fmt.Println(err)
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			err = json.Unmarshal(body, &ep)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(ep)
			episode := generateEpisodeString(ep)
			return episode
		}
		return "No Previous Episode"
	}
	if data.Links.Next != nil {
		resp, err := http.Get(data.Links.Next.Href)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		err = json.Unmarshal(body, &ep)
		if err != nil {
			fmt.Println(err)
		}
		episode := generateEpisodeString(ep)
		return episode
	}
	return "No Episode Scheduled"

}
