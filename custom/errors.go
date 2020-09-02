package main

const (
	errIncorrectSecret       = "incorrect secret received"
	errIncorrectRequestBody  = "incorrect request body received"
	errIncorrectHTTPMethod   = "unsupported http method received"
	errMarshallingCreateJSON = "error marshalling create json"
	errMarshallingUpdateJSON = "error marshalling update json"
	errCreatingAuth0Request  = "error creating auth0 request"
	errExecutingAuth0Request = "error executing auth0 request"
	errIncorrectResponseBody = "incorrect response body received"
	errDgraphMutation        = "error executing dgraph mutation"
)
