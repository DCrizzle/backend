# backend

## outline
- [ ] backend
	- [ ] update schema while running app
	- [ ] encrypt data received from api (not encrypted by api)
	- [ ] activity logging output
	- [ ] generate backups at specified frequency
		- [ ] note: https://medstack.co/blog/hipaa-tips-2-hipaa-compliant-databases/
	- [ ] handlers
		- [ ] "/org/{id}/graphql" -> execute mutations (post)
		- [ ] "/org/{id}/graphql" -> execute queries (get)
			- [ ] note: both endpoints receive all object CRUD operations
	- [ ] note: see heupr/core for csrf example
- [ ] frontend
	- [ ] endpoints
		- [ ] "/" -> redirect to "/login"
		- [ ] "/login" -> validate user login; redirect to "/org/{id}" (post)
		- [ ] "/org/{id}" -> render org (get)
		- [ ] "/org/{id}/db" -> execute mutations (post)
		- [ ] "/org/{id}/db" -> execute queries (get)
		- [ ] "/org/{id}/admin" -> org settings (get)
		- [ ] "/org/{id}/admin/users" -> users settings (get)
		- [ ] "/org/{id}/admin/user/{id}" -> user settings (get)

## notes
- [ ] `delete*` payload responses include a `msg` field that can be used as a default
- [ ] deleting any object from dgraph will remove all references to that object on other objects
- [ ] dgraph auto-generates crud operations and associated input/output types
- [ ] example authorization [directives](https://github.com/99designs/gqlgen/issues/785#issue-465696123)
- [ ] [graphql cheatsheet](https://devhints.io/graphql)
