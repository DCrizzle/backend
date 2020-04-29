# tbd

## outline
- [ ] backend
	- [ ] update schema while running app
	- [ ] encrypt data received from api (not encrypted by api)
	- [ ] activity logging output
	- [ ] generate backups at specified frequency
		- [ ] NOTE: https://medstack.co/blog/hipaa-tips-2-hipaa-compliant-databases/
	- [ ] handlers
		- [ ] "/org/{id}/graphql" -> execute mutations (POST)
		- [ ] "/org/{id}/graphql" -> execute queries (GET)
			- [ ] NOTE: both endpoints receive all object CRUD operations
	- [ ] NOTE: see heupr/core for CSRF example
- [ ] frontend
	- [ ] endpoints
		- [ ] "/" -> redirect to "/login"
		- [ ] "/login" -> validate user login; redirect to "/org/{id}" (POST)
		- [ ] "/org/{id}" -> render org (GET)
		- [ ] "/org/{id}/db" -> execute mutations (POST)
		- [ ] "/org/{id}/db" -> execute queries (GET)
		- [ ] "/org/{id}/admin" -> org settings (GET)
		- [ ] "/org/{id}/admin/users" -> users settings (GET)
		- [ ] "/org/{id}/admin/user/{id}" -> user settings (GET)

## notes
- [ ] `delete*` payload responses include a `msg` field that can be used as a default
- [ ] deleting any object from Dgraph will remove all references to that object on other objects
- [ ] dgraph auto-generates CRUD operations and associated input/output types
- [ ] example authorization [directives](https://github.com/99designs/gqlgen/issues/785#issue-465696123)
