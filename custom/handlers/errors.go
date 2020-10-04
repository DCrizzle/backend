package handlers

const (
	// ErrIncorrectSecret is an incorrect or missing secret value
	ErrIncorrectSecret = "incorrect secret received"
	// ErrIncorrectRequestBody is a malformed JSON body was received
	ErrIncorrectRequestBody = "incorrect request body received"
	// ErrIncorrectHTTPMethod is an unsupported HTTP method was received
	ErrIncorrectHTTPMethod = "unsupported http method received"
	// ErrMarshallingCreateJSON is an error marshalling create Auth0 user JSON
	ErrMarshallingCreateJSON = "error marshalling create user json"
	// ErrMarshallingUpdateJSON is an error marshalling update Auth0 user JSON
	ErrMarshallingUpdateJSON = "error marshalling update user json"
	// ErrCreatingAuth0Request is an error creating an Auth0 HTTP request
	ErrCreatingAuth0Request = "error creating auth0 user request"
	// ErrExecutingAuth0Request is an error executing an Auth0 HTTP request
	ErrExecutingAuth0Request = "error executing auth0 user request"
	// ErrUnmarshallingResponseBody is an error unmarshalling an Auth0 response
	ErrUnmarshallingResponseBody = "error unmarshalling auth0 response body"
	// ErrDgraphMutation is an error executing the custom server Dgraph mutation
	ErrDgraphMutation = "error executing dgraph mutation"
	// ErrMarshallingDgraphJSON is an error marshalling a Dgraph response
	ErrMarshallingDgraphJSON = "error marshalling dgraph response json"
	// ErrClassifyingEntities is an error invoking the classifier
	ErrClassifyingEntities = "error classifying entities"
)
