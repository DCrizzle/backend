package main

type dgraphRequest struct {
	Owner     *string `json:"owner,omitempty"`
	Auth0ID   *string `json:"authZeroID,omitempty"` // NOTE: "authZeroID" field name is necessary due to a Dgraph limitation
	Email     *string `json:"email,omitempty"`
	Password  *string `json:"password,omitempty"`
	Role      *string `json:"role,omitempty"`
	Org       *string `json:"org,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName  *string `json:"lastName,omitempty"`
}

type createUserRequest struct {
	Email       string      `json:"email" valid:"required"`
	Password    string      `json:"password" valid:"required"`
	AppMetadata appMetadata `json:"app_metadata" valid:"required"`
	FirstName   string      `json:"given_name" valid:"required"`
	LastName    string      `json:"family_name" valid:"required"`
	Connection  string      `json:"connection" valid:"required"`
}

type updateUserRequest struct {
	AppMetadata appMetadata `json:"app_metadata,omitempty"`
	Password    *string     `json:"password,omitempty"`
}

type appMetadata struct {
	Role  *string `json:"role,omitempty"`
	OrgID *string `json:"orgID,omitempty"` // NOTE: possibly change to "ownerID" for consistency
}

type auth0Response struct {
	Auth0ID string `json:"user_id"`
}
