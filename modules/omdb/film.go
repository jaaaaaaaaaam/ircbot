package omdb

import (
	"encoding/json"
	"fmt"
	"github.com/jaaaaaaaaaam/ircbot/modules/omdb/models"
	"github.com/kennygrant/sanitize"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var omdbAPI = "http://www.omdbapi.com/?t=%s&y=&plot=short&r=json"

func generateString(film models.Film) string {
	str := os.Getenv("FILMSTRING")
	r := strings.NewReplacer(
		"#title#", film.Title,
		"#year#", film.Year,
		"#rating#", film.Rated,
		"#release#", film.Released,
		"#runtime#", film.Runtime,
		"#genre#", film.Genre,
		"#director#", film.Director,
		"#writer#", film.Writer,
		"#actors#", film.Actors,
		"#plot#", film.Plot,
		"#language#", film.Language,
		"#country#", film.Country,
		"#awards#", film.Awards,
		"#posterURL#", film.Poster,
		"#metascore#", film.Metascore,
		"#imdb-rating#", film.ImdbRating, //fmt.Sprintf("%f", film.ImdbRating),
		"#imdb-votes#", film.ImdbVotes,
		"#imdb-ID#", fmt.Sprintf("http://www.imdb.com/title/%s/", film.ImdbID),
		"#type#", film.Type,
	)
	result := r.Replace(str)
	ret, err := strconv.Unquote(`"` + result + `"`)
	if err != nil {
		fmt.Println("error replacing")
		fmt.Println(err)
	}
	return sanitize.HTML(ret)
}

// FilmLookup finds the film from omdb
func FilmLookup(search string) string {
	var film models.Film
	q := fmt.Sprintf(omdbAPI, search)
	resp, err := http.Get(q)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if string(body) == "" {
		return "This show doesn't exist!"
	}

	err = json.Unmarshal(body, &film)
	if err != nil {
		fmt.Println(err)
	}
	returnString := generateString(film)
	return returnString
}
