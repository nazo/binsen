[![CircleCI](https://circleci.com/gh/nazo/binsen/tree/master.svg?style=svg)](https://circleci.com/gh/nazo/binsen/tree/master)
[![codecov](https://codecov.io/gh/nazo/binsen/branch/master/graph/badge.svg)](https://codecov.io/gh/nazo/binsen)

# binsen - Shared Markdown Document Editor

This project is WIP, DO NOT USE PRODCUTION.

# development

- cp ./server/.env.example ./server/.env
- edit `./server/.env`
- run `docker-compose up`
- run `cd server && make migrate && cd ..`
- run `docker-compose exec server go run cmd/main.go u add [your gmail address]`

# production

- WIP

# software stack

## Frontend

- typescript
- [vue.js](https://github.com/vuejs/vue)
- [nuxt.js](https://github.com/nuxt/nuxt.js)
- [jest](https://github.com/facebook/jest)

## Backend

- go
- [echo](https://github.com/labstack/echo)
- [sqlboiler](https://github.com/volatiletech/sqlboiler)
