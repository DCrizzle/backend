package users

const (
	errorIncorrectRequestBody  = "incorrect request body received"
	errorIncorrectHTTPMethod   = "unsupported http method received"
	errorCreateAuth0User       = "error creating auth0 user"
	errorUpdateAuth0User       = "error updating auth0 user"
	errorDeleteAuth0User       = "error deleting auth0 user"
	errorDgraphMutation        = "error executing dgraph mutation"
	errorMarshallingDgraphJSON = "error marshalling dgraph response json"
)
