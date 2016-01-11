# ircbot

This is a small project for me to learn Go.

It uses the TVMaze API and also OMDB for film information.

The current commands for the bot are...

`!show <search>` e.g. `!show brooklyn nine nine` _also has an alias_ `!tv`

`!film <search>` e.g. `!film gladiator` _also has an alias_ `!movie`

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
FILMSTRING=
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
FILMSTRING="#title# - #plot# - #runtime# - #genre# - #imdb-rating#/10 - #imdb-ID#"
HUMANIZE=1
```

This will output the following:
```
<jaaaaaaaaaam>    !show gotham
<Awesomebot>      Looking up 'gotham'
<Awesomebot>      Gotham | Next: 1 month from now - 2x12 - Rise of the Villains: Cold, Dark Night | Prev: 1 month ago - 2x11 - Rise of the Villains: Worse Than a Crime | Premiered: 2014-09-22
```
##### NICK

This is the nick of the bot.

`NICK=Steve`


##### USERNAME

This is the username of the bot.

`USERNAME=Robot`

These options would result in the following:

`Steve!Robot@hostname`

```
<Steve> Bleep Bloop Bloop
<Steve> I'm an IRC Bot written in Go!
```

##### CHANS

The channel you would like the bot to join.

`CHAN="#spamroom"`

##### Network

The IRC Network and port you would like the bot to connect to.

`irc.network.com:6667`

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

##### FILMSTRING

This is used to format the output of the !imdb command

| Option | Description |
| --------- | --------------- |
| #title# | Film title |
| #year# | TVMaze episode URL |
| #rating# | The MPAA film rating |
| #release# | The release date of the film |
| #runtime# | The films runtime |
| #genre# | The genre of the film |
| #director# | The film's directors |
| #writer# | The films writers |
| #actors# | The films actors |
| #plot# | Short plot of the film |
| #language# | The films language |
| #country# | The films countries |
| #awards# | Awards the film has |
| #posterURL# | URL for the film poster |
| #metascore | The films metascore |
| #imdb-rating# | The films IMDB rating |
| #imdb-votes# | The films IMDB votes |
| #imdb-ID# | The url to the IMDB page for the film |
| #type# | The type of the search query (Film etc) |

##### HUMANIZE

This option allows the dates of next/previous episodes to be formatted into human readable strings.

There are currently 3 options.

```
0 - Turns off date humanizations
1 - Humanization using github.com/dustin/go-humanize (1 month ago, in 1 year) doesn't account for overages, 1 month and 15 days is displayed as 1 month.
2 - Humanization into days only. (2 days ago/in 125 days)
```

### TODO
- Bot Basics
  - [ ] Join channels based on comma separated list in .env
  - [ ] Nickserv operations
  - [ ] .env option for greeting other users
- TVMaze modules
  - [x] Finish Show struct for JSON data from TVMaze
  - [x] Show desired output from show
  - [x] Finish adding options for output
  - [x] Get the dates/episode info for previous/next episodes
  - [x] Show the dates in human readable format e.g. (2 days ago, in 3 months, 2 days etc)
  - [ ] Allow users to choose a format for date's if humanization is disabled
  - [ ] Show the user suggestions if their search returns nothing
- OMDB module
  - [ ] Filter out not found requests and show an error instead
