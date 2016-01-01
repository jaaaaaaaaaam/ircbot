package tvmaze

import (
	"github.com/jaaaaaaaaaam/ircbot/modules/tvmaze/models"
)

// ProcessEpisode processes either the next/prev episode data
func ProcessEpisode(data models.ShowLinks, order string) string {
	if order == "prev" {
		if data.Previous != nil {

		}
		return "No Previous Episode"
	}
	if data.Next != nil {

	}
	return "No Episode Scheduled"

}
