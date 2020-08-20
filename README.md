# backend

## docker

## dgraph

- [ ] `delete*` payload responses include a `msg` field that can be used as a default
- [ ] deleting any object from dgraph will remove all references to that object on other objects
- [ ] [graphql cheatsheet](https://devhints.io/graphql)
- [ ] include an "attribute mapping" type on the Lab entities (?)
- - [ ] possibly to show individual Lab header/field names for attributes within the "root" type
- - [ ] this would be added/updated as Lab nodes are added/updated
- [ ] general notes:
- - [ ] update relationships using "set"/"remove" keywords in queries
- - [ ] searching: https://dgraph.io/docs/query-language/#indexing
- - [ ] "query*"
- - - [ ] "filter" field contains fields tagged with "@search" directive

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
	- [ ] build command: `go build -tags mock`
- [ ] testing - running isolated for code testing purposes (e.g. "go test")
- [ ] docker:
	- [ ] build with `docker build .`
	- [ ] copy result output ID
	- [ ] run with `docker run -p 4080:4080 <ID>`

## notes

- [ ] ports
	- [ ] dgraph zero: 5080/6080
	- [ ] dgraph alpha: 7080/8080/9080
	- [ ] helper: 4080
