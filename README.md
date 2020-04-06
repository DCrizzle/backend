# tbd

## outline
- [ ] api
	- [ ] update schema while running app
	- [ ] encrypt data received from api (not encrypted by api)
	- [ ] activity logging output
	- [ ] generate backups at specified frequency
		- [ ] NOTE: https://medstack.co/blog/hipaa-tips-2-hipaa-compliant-databases/
	- [ ] graph queries (and general flow)
		- [ ] fixed: predetermined, for organizations/users, fetching schema
		- [ ] flexible: database query filter, changes based on schema fetches
	- [ ] handlers
		- [ ] "/org/{id}/graphql" -> execute mutations (POST)
		- [ ] "/org/{id}/graphql" -> execute queries (GET)
			- [ ] NOTE: both endpoints receive all object CRUD operations
	- [ ] NOTE: see heupr/core for CSRF example
- [ ] ui
	- [ ] user types:
		- [ ] member (owner/admin, "storage companies")
		- [ ] visitor (user/searcher, "laboratory companies")
	- [ ] architecture:
		- [ ] "/" -> public landing page
		- [ ] "/login" -> query backend w/ provided info/tokens, redirect to "/org/{id}"
		- [ ] "/org/{id}" -> query org contents
		- [ ] "/org/{id}/db/query" -> query schema, render filter page
		- [ ] "/org/{id}/db/mutate" -> input new data (?)
		- [ ] "/org/{id}/admin" -> query org, general org settings (maybe under "/settings")
		- [ ] "/org/{id}/admin/users" -> query/list org users, send invites
		- [ ] "/org/{id}/admin/users/{id}" -> query single user
