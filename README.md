# backend

## general

Several prerequisites are needed in order to run the `backend` package locally. Follow the installation/download instructions for your local operating system.

- install [Git](https://git-scm.com/downloads)
	- [additional installation information](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- install [Go](https://golang.org/doc/install)
- install [Postman](https://www.postman.com/downloads/)
- install [Dgraph](https://github.com/dgraph-io/dgraph#install-from-source)
	- run `git clone git@github.com:dgraph-io/dgraph.git` in the command line
	- run the **Install from Source** instructions in the link above

**NOTE**: the `master` branch version of Dgraph is being used to accommodate pre-release features

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

To interact with the Dgraph database and schema, follow the steps below.

- launch **Dgraph** locally
	- run `dgraph zero --my=localhost:5080` in a command line tab
	- run `dgraph alpha --lru_mb=2048 --my=localhost:7080` in another command line tab
	- in another command line tab, navigate into the root directory of where you've installed the `backend` repository
	- run `curl -X POST localhost:8080/admin/schema --data-binary '@database/schema.graphql'`
- build and run the **`demo`** package
	- run `cd bin/demo` to navigate into the `demo` package directory
	- run `go run main.go` to begin loading demo data into the running Dgraph database
		- if an error saying `"401 status received - management api token may be expired"` the Auth0 Management API Token needs to be updated
		- go to the Auth0 Dashboard
		- click **APIs** -> **Auth0 Management API** -> **Test**
		- from the **Response** field, copy the value of `"access_token"`, starting with `"ey..."`
		- open the `backend/etc/config/config.json` file
		- paste the token value into the `AUTH0.MANAGEMENT_TOKEN` field
	- a progress bar will display as data is loaded into Dgraph - once complete the database may be used
- setup and use **Postman**
	- launch Postman
	- click **Import** -> **Choose Files**
	- find, select, and upload the `backend/etc/postman/collection.json` file
	- run `cd ../token` in the command line
	- run `go run main.go`
	- copy the output token value
	- in your Postman requests, under the **Headers** tab, add a `Key` with the value "`Authorization`" and a `Value` with a value of the copied token
	- Postman requests are now authorized to communicate with the Dgraph database

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

- [ ] ports
	- [ ] dgraph zero: 5080/6080
	- [ ] dgraph alpha: 7080/8080/9080
	- [ ] helper: 4080
