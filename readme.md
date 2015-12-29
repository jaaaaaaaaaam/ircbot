# ircbot
IRC Bot created in GO.

## How to run

`git clone https://github.com/jaaaaaaaaaam/ircbot.git`

#### Set up env file
```
cp .env.example .env
vi .env
```

`go run cli/app/main.go`

### TODO
- Bot Basics
  - [ ] Join channels based on comma separated list in .env
- TVMaze modules
  - [x] Finish Show struct for JSON data from TVMaze
  - [x] Show desired output from show
  - [ ] Finish adding options for output
  - [ ] Show the user suggestions if their search returns nothing
