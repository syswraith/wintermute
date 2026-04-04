# Wintermute

url shortener

## What's done so far

- [x] gin setup
- [x] mysql setup
- [x] gorm setup
- [x] create url logic
- [x] fetch url logic
- [x] create url frontend
- [x] expiry
- [x] auth (jwt)

## Setup

### Required software

- go
- bun
- mysql

### Setup instructions

```bash
git clone https://github.com/syswraith/wintermute
cd wintermute

# venus     = frontend
# vulcan    = backend
# minerva   = db

# mysql setup
# create database wintermute;
# username = "user"
# password = "password"
# gorm will do the rest

# start vulcan
cd vulcan
go mod tidy     # install dependencies
go run .        # additionally run with air for hot reload

# start venus
cd venus
bun install     # install dependencies
bun run dev     # --hot for hot reload
```
