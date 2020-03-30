# tbd

## outline
- [ ] app
	- [ ] update schema while running app
		- [ ] NOTE: may not really be needed since api/methods would need to be updated
		- [ ] variable in schema.go
	- [x] encrypt data received from api (not encrypted by api)
	- [ ] obfuscate dgraph/graphql implementation
		- [ ] include user/session/datetime/etc for operations
		- [ ] NOTE: possibly restructure to interface (operation) w/ sub interface (mutation/query)
		- [ ] NOTE: unsure on this; app might be encrypted/secure dgraph
	- [ ] activity logging output
	- [ ] generate backups at specified frequency
	- [ ] NOTE: https://medstack.co/blog/hipaa-tips-2-hipaa-compliant-databases/
- [ ] api
	- [ ] objects
		- [ ] org(s)
		- [ ] user(s)
- [ ] ui
	- [ ] storage (storage companies)
		- [ ] specimen
			- [ ] add multi
			- [ ] add single
			- [ ] delete multi
			- [ ] delete single
			- [ ] list all
			- [ ] list filter
	- [ ] search (laboratories/researchers)
