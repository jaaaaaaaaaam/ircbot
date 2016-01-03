package tvmaze

import (
	"encoding/json"
	"fmt"
	"github.com/jaaaaaaaaaam/ircbot/modules/tvmaze/models"
	//"github.com/jinzhu/now"
	"github.com/dustin/go-humanize"
	"github.com/kennygrant/sanitize"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func generateEpisodeString(ep models.ShowEpisode) string {
	str := os.Getenv("EPSTRING")
	human := os.Getenv("HUMANIZE")
	airdate := ""
	if human == "1" {
		time, err := time.Parse("2006-01-02", ep.Airdate)
		if err != nil {
			fmt.Println(err)
		}
		airdate = humanize.Time(time)
	} else {
		airdate = ep.Airdate
	}

	r := strings.NewReplacer(
		"#ID#", string(ep.ID),
		"#URL#", ep.URL,
		"#name#", ep.Name,
		"#season#", strconv.Itoa(ep.Season),
		"#episode#", strconv.Itoa(ep.Number),
		"#airdate#", airdate, //time.Parse("2000-01-01", ep.Airdate),
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
	var ep models.ShowEpisode
	if order == "prev" {
		if data.Links.Previous != nil {
			epjson, err := http.Get(data.Links.Previous.Href)
			if err != nil {
				fmt.Println(err)
			}
			defer epjson.Body.Close()
			body, err := ioutil.ReadAll(epjson.Body)

			err = json.Unmarshal(body, &ep)
			if err != nil {
				fmt.Println(err)
			}
			episode := generateEpisodeString(ep)
			return episode
		}
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
	return "Something went seriously wrong"
}
