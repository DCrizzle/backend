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
