package demo

const (
	specimenCount     = 1
	consentFormCount  = 1
	ownerOrgCount     = 1
	labOrgCount       = 1
	storageOrgCount   = 1
	protocolCount     = 1
	protocolFormCount = 1
	testCount         = 1
)

type donor struct {
	Street    string   `json:"street"`
	City      string   `json:"city"`
	County    string   `json:"county"`
	State     string   `json:"state"`
	ZIP       int      `json:"zip"`
	Owner     string   `json:"owner"`
	Age       int      `json:"age"`
	DOB       string   `json:"dob"`
	Sex       string   `json:"sex"`
	Race      string   `json:"race"`
	Specimens []string `json:"specimens"`
	Consents  []string `json:"consents"`
}

type consent struct {
	Owner           string `json:"owner"`
	Donor           string `json:"donor"`
	Specimen        string `json:"donor"`
	Protocol        string `json:"donor"`
	ConsentDate     string `json:"consentedDate"`
	DestructionDate string `json:"destructionDate"`
}

type consentForm struct {
	Owner    string   `json:"owner"`
	Title    string   `json:"title"`
	Body     string   `json:"body"`
	Consents []string `json:"consents"`
}

type org struct {
	Street    string   `json:"street"`
	City      string   `json:"city"`
	County    string   `json:"county"`
	State     string   `json:"state"`
	ZIP       int      `json:"zip"`
	Name      string   `json:"name"`
	Users     []string `json:"users"`
	CreatedOn string   `json:"createdOn"`
	UpdatedOn string   `json:"updatedOn"`
}

type ownerOrg struct {
	org
	Labs     []string `json:"labs"`
	Storages []string `json:"storages"`
}

type labOrg struct {
	org
	Owner     string   `json:"owner"`
	Specimens []string `json:"specimens"`
	Plans     []string `json:"plans"`
}

type storageOrg struct {
	org
	Owner     string   `json:"owner"`
	Specimens []string `json:"specimens"`
	Plans     []string `json:"plans"`
}

type protocol struct {
	Street      string   `json:"street"`
	City        string   `json:"city"`
	County      string   `json:"county"`
	State       string   `json:"state"`
	ZIP         int      `json:"zip"`
	Owner       string   `json:"owner"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Form        string   `json:"form"`
	Plan        string   `json:"plan"`
	Ages        []int    `json:"ages"`
	AgeStart    int      `json:"ageStart"`
	AgeEnd      int      `json:"ageEnd"`
	DOBs        []string `json:"dobs"`
	DOBStart    string   `json:"dobStart"`
	DOBEnd      string   `json:"dobEnd"`
	Race        string   `json:"race"`
	Sex         string   `json:"sex"`
	Specimens   []string `json:"specimens"`
}

type protocolForm struct {
	Owner      string   `json:"owner"`
	Title      string   `json:"title"`
	Body       string   `json:"body"`
	ProtocolID string   `json:"protocolID"`
	Protocols  []string `json:"protocols"`
}

type tests struct {
	Description string   `json:"description"`
	Owner       string   `json:"owner"`
	Labs        []string `json:"labs"`
	Specimens   []string `json:"specimens"`
	Results     []string `json:"results"`
}

type bloodSpecimen struct {
	ExternalID      string   `json:"externalID"`
	Type            string   `json:"type"`
	CollectionDate  string   `json:"collectionDate"`
	Donor           string   `json:"donor"`
	Container       string   `json:"container"`
	Status          string   `json:"status"`
	DestructionDate string   `json:"destructionDate"`
	Description     string   `json:"description"`
	Consent         string   `json:"consent"`
	Owner           string   `json:"owner"`
	Lab             string   `json:"lab"`
	Storage         string   `json:"storage"`
	Protocol        string   `json:"protocol"`
	Tests           []string `json:"tests"`
	BloodType       string   `json:"bloodType"`
	Volume          float64  `json:"volume"`
}
