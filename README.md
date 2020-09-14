# backend

> The API and database :sloth:

## general

Several prerequisites are needed in order to run the `backend` package locally. Follow the installation/download instructions for your local operating system.

- install [Git](https://git-scm.com/downloads)
	- [additional installation information](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- install [Go](https://golang.org/doc/install)
- install [Postman](https://www.postman.com/downloads/)
- install [Dgraph](https://github.com/dgraph-io/dgraph#install-from-source)
	- run `git clone git@github.com:dgraph-io/dgraph.git` in the command line
	- run the **Install from Source** instructions in the link above

These should be the minimum resources needed to get up and running with `backend`! :thumbsup:

**NOTE**: the `master` branch version of Dgraph is being used to accommodate pre-release features

## docker

- [ ] tbd

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

To interact with the Dgraph database and schema, follow the steps below with all commands being run in the command line from the root of the `backend` repository.

1. run `chmod -R +x ./bin/`
2. launch **Dgraph** locally by running `./bin/start_dgraph`
3. build and run the **`demo`** package by running `./bin/load_demo`; a progress bar will display as data is loaded into Dgraph and once complete the database may be used
4. launch **Postman** and click **Import** -> **Choose Files** then find, select, and upload the `backend/etc/postman/collection.json` file
	- run `./bin/get_token` from the root of the `backend` repository and copy the output token value
	- in each of your Postman requests, under the **Headers** tab, add a `Key` with the value "`X-Auth0-Token`" and a `Value` with a value of the copied token and the requests are now authorized to communicate with the Dgraph database

**NOTE**: the Management API key fetched from Auth0 is **_sensitive data_** and should not be shared publicly  
**NOTE**: currently the user JWT issued from the `token` package is configured to `john.forstmeier@gmail.com` in Auth0  

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

#### Ports currently used

- Dgraph Zero: 5080 and 6080
- Dgraph Alpha: 7080, 8080, and 9080
- custom server: 4080
