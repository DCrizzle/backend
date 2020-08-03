# backend

## docker

## dgraph

- [ ] `delete*` payload responses include a `msg` field that can be used as a default
- [ ] deleting any object from dgraph will remove all references to that object on other objects
- [ ] [graphql cheatsheet](https://devhints.io/graphql)

## demo

- [ ] outline:
- [ ] launch `frontend` package
- [ ] open up console on "network" tab
- [ ] log in to app
- [ ] retrieve jwt from "token" row values
- [ ] pass jwt into X-Auth0-Token header

## helper

- [ ] outline:
- [ ] holds server for intercepting dgraph @custom requests
- [ ] responsible for enriching/handling requests
- [ ] also servers as a way to mock/isolate dgraph @custom directives
- [ ] available configurations/states:
- [ ] production - running fully connected to external resources (e.g. auth0)
- [ ] mocking - running dgraph live but not unit testing (e.g. manually testing database/api)
- [ ] testing - running isolated for code testing purposes (e.g. "go test")
