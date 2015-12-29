package tvmaze

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

var showAPI = "http://api.tvmaze.com/singlesearch/shows?q="

// Show is the main JSON object
type Show struct {
	ID      int      `json:"id"`
	URL     string   `json:"url"`
	Name    string   `json:"name"`
	Type    string   `json:"type"`
	Genres  []string `json:"genres"`
	Status  string   `json:"status"`
	Runtime int      `json:"runtime"`
}

// ShowLookup looks for the show and then returns a message
func ShowLookup(search string) string {
	var show Show
	q := strings.Join([]string{showAPI, search}, "")
	//request, err := http.NewRequest("GET", q)
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println("no")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//err := json.Unmarshal([]byte(string(body)), $show)
	er := json.Unmarshal([]byte(string(body)), &show)
	if er != nil {
		panic(er)
	}
	fmt.Println(show.Genres[1])
	return strconv.Itoa(show.ID)
}
