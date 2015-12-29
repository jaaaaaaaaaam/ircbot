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
#showname# - Name of the show
#status# - Status of the show (continuing/cancelled/ended etc)
#network.name# - The name of the network that the show airs on
```


### TODO
- Bot Basics
  - [ ] Join channels based on comma separated list in .env
- TVMaze modules
  - [x] Finish Show struct for JSON data from TVMaze
  - [x] Show desired output from show
  - [ ] Finish adding options for output
  - [ ] Show the user suggestions if their search returns nothing
