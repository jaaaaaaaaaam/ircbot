# ircbot
IRC Bot created in GO.

## How to run

`git clone https://github.com/jaaaaaaaaaam/ircbot.git`

### Set up env file
```
cp .env.example .env
vi .env
```

`go run cli/app/main.go`

### .env options

```
NICK=
USERNAME=
CHANS=
NETWORK=
SHOWSTRING=
EPSTRING=
HUMANIZE=
```

_example_
```
NICK=Awesomebot
USERNAME=John
CHANS="#botchan"
NETWORK=irc.network.net:6667
SHOWSTRING="\u0002#showname#\u0002 | Next: #nextEp# | Prev: #previousEp# | Premiered: #premiered#"
EPSTRING="#airdate# - #season#x#episode# - #name#"
HUMANIZE=1
```

This will output the following:
```
<jaaaaaaaaaam>    !show gotham
<jambot>          Looking up 'gotham'
<jambot>          Gotham | Next: 1 month from now - 2x12 - Rise of the Villains: Cold, Dark Night | Prev: 1 month ago - 2x11 - Rise of the Villains: Worse Than a Crime | Premiered: 2014-09-22
```


##### SHOWSTRING

This is used to format the output for !show &lt;showname&gt;

###### Current options:
| Option | Description |
| --------- | --------------- |
| #ID# | TVMaze ID for the show |
| #URL# | TVMaze url for the show |
| #showname# | Name of the show |
| #type# | The type of show |
| #genres# | Comma separated list of genres for the show |
| #status# | Status of the show (continuing/cancelled/ended etc) |
| #runtime# | The runtime of the show in minutes |
| #premiered# | The date the show was premiered |
| #schedule.time# | The time that the show usually airs |
| #schedule.days# | Comma separated list of days the show airs on |
| #rating# | The average rating of the show |
| #network.name# | The name of the network that the show airs on |
| #network.country.name# | The name of the country the network is in |
| #network.country.code# | The country code of the country the network is in |
| #network.country.timezone# | The timezone that the network is in |
| #webchannel# | The name of the webchannel the show airs on |
| #externals.tvrage# | The 'TVRage' ID for the show |
| #externals.thetvdb# | The 'TheTVDB' ID for the show |
| #image.medium# | Medium sized image for the show |
| #image.original# | Original sized image for the show |
| #summary# | The show summary |
| #updated# | Unix timestamp of when the info was last updated |
| #links.self# | Link to API for the current set of data |
| #previousEp# | Formatted data for the previous episode, see EPSTRING below |
| #nextEp# | Formatted data for the next episode, see EPSTRING below |

##### EPSTRING

This is used to format the output of the information for the next episode or previous episode.

| Option | Description |
| --------- | --------------- |
| #ID# | TVMaze episode ID |
| #URL# | TVMaze episode URL |
| #name# | Name of the episode |
| #season# | The season number |
| #episode# | The episode number |
| #airdate# | The episode's airdate |
| #airtime# | The episode's airtime |
| #airstamp# | The full datetime of the episode's airing |
| #runtime# | The runtime of the episode in minutes |


### TODO
- Bot Basics
  - [ ] Join channels based on comma separated list in .env
  - [ ] Nickserv operations
- TVMaze modules
  - [x] Finish Show struct for JSON data from TVMaze
  - [x] Show desired output from show
  - [x] Finish adding options for output
  - [x] Get the dates/episode info for previous/next episodes
  - [x] Show the dates in human readable format e.g. (2 days ago, in 3 months, 2 days etc)
  - [ ] Show the user suggestions if their search returns nothing
