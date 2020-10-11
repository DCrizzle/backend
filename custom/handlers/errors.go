package handlers

const (
	// ErrIncorrectSecret is an incorrect or missing secret value.
	ErrIncorrectSecret = "incorrect secret received"
	// ErrIncorrectRequestBody is a malformed JSON body was received.
	ErrIncorrectRequestBody = "incorrect request body received"
	// ErrIncorrectHTTPMethod is an unsupported HTTP method was received.
	ErrIncorrectHTTPMethod = "unsupported http method received"
	// ErrUnmarshallingResponseBody is an error unmarshalling an Auth0 response.
	ErrUnmarshallingResponseBody = "error unmarshalling auth0 response body"
	// ErrDgraphMutation is an error executing the custom server Dgraph mutation.
	ErrDgraphMutation = "error executing dgraph mutation"
	// ErrMarshallingDgraphJSON is an error marshalling a Dgraph response.
	ErrMarshallingDgraphJSON = "error marshalling dgraph response json"
	// ErrClassifyingEntities is an error invoking the classifier.
	ErrClassifyingEntities = "error classifying entities"
)
