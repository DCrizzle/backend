package handlers

// DgraphRequest is a User type @custom directive request from Dgraph
// sent to the custom server
type DgraphRequest struct {
	Owner     *string `json:"owner,omitempty"`
	Auth0ID   *string `json:"authZeroID,omitempty"` // NOTE: "authZeroID" field name is necessary due to a Dgraph limitation
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
	Role      *string `json:"role,omitempty"`
	Org       *string `json:"org,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

// CreateUserRequest is a User type @custom directive requset from the
// custom server to Auth0 to create an Auth0 user
type CreateUserRequest struct {
	Email       string      `json:"email" valid:"required"`
	Password    string      `json:"password" valid:"required"`
	AppMetadata AppMetadata `json:"app_metadata" valid:"required"`
	FirstName   string      `json:"given_name" valid:"required"`
	LastName    string      `json:"family_name" valid:"required"`
	Connection  string      `json:"connection" valid:"required"`
}

// UpdateUserRequest is a User type @custom directive requset from the
// custom server to Auth0 to update the Auth0 user
type UpdateUserRequest struct {
	AppMetadata AppMetadata `json:"app_metadata,omitempty"`
	Password    *string     `json:"password,omitempty"`
}

// AppMetadata is the app_metadata object on the user in Auth0
type AppMetadata struct {
	Role  *string `json:"role,omitempty"`
	OrgID *string `json:"orgID,omitempty"` // NOTE: possibly change to "ownerID" for consistency
}

// Auth0Response is the response from Auth0 to the custom server request
type Auth0Response struct {
	Auth0ID string `json:"user_id"`
}

// DgraphEntitiesRequest is a request to identify entities in a text blob
type DgraphEntitiesRequest struct {
	Owner   string `json:"owner"`
	Form    string `json:"form"`
	DocType string `json:"docType"`
	Blob    string `json:"blob"`
}

// EntitiesResponse is a response from entities classification
type EntitiesResponse struct {
	Data map[string][]string `json:"data"`
}
