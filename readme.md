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

##### SHOWSTRING

This is used to format the output for !show &lt;showname&gt;

###### Current options:
```
#ID# - TVMaze ID for the show
#URL# - TVMaze url for the show
#showname# - Name of the show
#type# - The type of show
#genres# - Comma separated list of genres for the show
#status# - Status of the show (continuing/cancelled/ended etc)
#runtime# - The runtime of the show in minutes
#premiered# - The date the show was premiered
#schedule.time# - The time that the show usually airs
#schedule.days# - Comma separated list of days the show airs on
#rating# - The average rating of the show
#network.name# - The name of the network that the show airs on
#network.country.name# - The name of the country the network is in
#network.country.code# - The country code of the country the network is in
#network.country.timezone# - The timezone that the network is in
#webchannel# - The name of the webchannel the show airs on
#externals.tvrage# - The 'TVRage' ID for the show
#externals.thetvdb# - The 'TheTVDB' ID for the show
#image.medium# - Medium sized image for the show
#image.original# - Original sized image for the show
#summary# - The show summary
#updated# - Unix timestamp of when the info was last updated
#links.self# - Link to API for the current set of data
#links.previous# - Link to the API for the previous episode
#links.next# - Link to the API for the next episode
```


### TODO
- Bot Basics
  - [ ] Join channels based on comma separated list in .env
  - [ ] Nickserv operations
- TVMaze modules
  - [x] Finish Show struct for JSON data from TVMaze
  - [x] Show desired output from show
  - [x] Finish adding options for output
  - [ ] Show the user suggestions if their search returns nothing
