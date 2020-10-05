# backend

> The API and database :sloth:

## general

Several prerequisites are needed in order to run the `backend` package locally. Follow the installation/download instructions for your local operating system.

- install [Git](https://git-scm.com/downloads)
	- [additional installation information](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- install [Go](https://golang.org/doc/install)
- install [Insomnia](https://insomnia.rest/download/core/?&ref=https%3A%2F%2Fgraphql.dgraph.io%2Fdocs%2Fquick-start%2F)
- install [Dgraph](https://github.com/dgraph-io/dgraph#install-from-source)
	- run `git clone git@github.com:dgraph-io/dgraph.git` in the command line
	- run the **Install from Source** instructions in the link above

These should be the minimum resources needed to get up and running with `backend`! :thumbsup:

**NOTE**: the `master` branch version of Dgraph is being used to accommodate pre-release features

## packages

### custom

`custom` intercepts and processes all Dgraph `@custom` directive requests. This is to allow the additional preprocessing ahead of the main GraphQL logic and to provide a base for local mocking. **No "smarts"** will be built into this package and it will _only be an intermediary_ responsible for invoking external APIs (e.g. Auth0) and internal packages (e.g. `entities`) to fulfill the required logic. Most or all of the internal packages will be housed in the private `internal` repository.

The server can be built in two configurations:

- production: `go build`
- mock: `go build -tags mock`

In both cases calling `./bin/start_custom` would start the server.

### subscription (tbd)

- intercepts `@withSubscription` directives
- configures/establishes GraphQL subscription link
- calls external APIs/internal packages as needed
- exposes webhook listener configuration endpoints
- manages/stores webhook listeners in database (schema likely requires update)
- emits received GraphQL subscription events on webhook listeners

### demo

`demo` is responsible for loading demo data into the Dgraph database. Follow the instructions below to get setup and execute all commands in the terminal from the root of the `backend` repository.

1. run `chmod -R +x ./bin/`
2. launch **Dgraph** locally by running `./bin/start_dgraph`
3. build and run the **`demo`** package by running `./bin/load_demo`; a progress bar will display as data is loaded into Dgraph and once complete the database may be used
4. launch **Insomnia** and contact the repo maintainer for the requests collection file
	- run `./bin/get_token` from the root of the `backend` repository and copy the output token value
	- in each of your Insomnia requests, under the **Headers** tab, add a `Key` with the value "`X-Auth0-Token`" and a `Value` with a value of the copied token and the requests are now authorized to communicate with the Dgraph database

**NOTE**: the Management API key fetched from Auth0 is **_sensitive data_** and should not be shared publicly  
**NOTE**: currently the user JWT issued from the `token` package is configured to `john.forstmeier@gmail.com` in Auth0  

## notes

### ports

- Dgraph Zero: `5080` and `6080`
- Dgraph Alpha: `7080`, `8080`, and `9080`
- custom server: `4080`

### design

- `struct` should be used for all request/response objects - specific `struct` objects with appropriate fields and JSON tags should be defined
- `map` should be used for Dgraph GraphQL mutation variables - a more flexible `map[string]interface{}` can be used when processing mutations
