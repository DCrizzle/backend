package graphql

const (
	// AddOwnerOrgsMutation is the GraphQL mutation that adds Org type objects to the Dgraph database.
	AddOwnerOrgsMutation = "mutation AddOwnerOrgs($input: [AddOwnerOrgInput!]!) { addOwnerOrg(input: $input) { ownerOrg { id } } }"

	// AddLabOrgsMutation is the GraphQL mutation that adds Lab type objects to the Dgraph database.
	AddLabOrgsMutation = "mutation AddLabOrgs($input: [AddLabOrgInput!]!) { addLabOrg(input: $input) { labOrg { id } } }"

	// AddStorageOrgsMutation is the GraphQL mutation that adds Storage type objects to the Dgraph database.
	AddStorageOrgsMutation = "mutation AddStorageOrgs($input: [AddStorageOrgInput!]!) { addStorageOrg(input: $input) { storageOrg { id } } }"

	// AddUsersMutation is the GraphQL mutation that adds User type objects to the Dgraph database.
	AddUsersMutation = "mutation AddUsers($input: [AddUserInput!]!) { addUser(input: $input) { user { email } } }"

	// UpdateUserMutation is the GraphQL mutation that updates User type objects to the Dgraph database.
	UpdateUserMutation = "mutation EditUser($input: UpdateUserInput!) { updateUser(input: $input) { user { email } } }"

	// DeleteUserMutation is the GraphQL mutation that deletes User type objects to the Dgraph database.
	DeleteUserMutation = "mutation RemoveUser($filter: UserFilter!) { deleteUser(filter: $filter) { user { email } } }"

	// AddProtocolsMutation is the GraphQL mutation that adds Protocol type objects to the Dgraph database.
	AddProtocolsMutation = "mutation AddProtocols($input: [AddProtocolInput!]!) { addProtocol(input: $input) { protocol { id } } }"

	// AddProtocolFormsMutation is the GraphQL mutation that adds ProtocolForm type objects to the Dgraph database.
	AddProtocolFormsMutation = "mutation AddProtocolForms($input: [AddProtocolFormInput!]!) { addProtocolForm(input: $input) { protocolForm { id } } }"

	// AddPlansMutation is the GraphQL mutation that adds Plan type objects to the Dgraph database.
	AddPlansMutation = "mutation AddPlans($input: [AddPlanInput!]!) { addPlan(input: $input) { plan { id } } }"

	// AddConsentFormsMutation is the GraphQL mutation that adds ConsentForm type objects to the Dgraph database.
	AddConsentFormsMutation = "mutation AddConsentForms($input: [AddConsentFormInput!]!) { addConsentForm(input: $input) { consentForm { id } } }"

	// AddDonorsMutation is the GraphQL mutation that adds Donor type objects to the Dgraph database.
	AddDonorsMutation = "mutation AddDonors($input: [AddDonorInput!]!) { addDonor(input: $input) { donor { id } } }"

	// AddConsentsMutation is the GraphQL mutation that adds Consetn type objects to the Dgraph database.
	AddConsentsMutation = "mutation AddConsent($input: [AddConsentInput!]!) { addConsent(input: $input) { consent { id } } }"

	// AddBloodSpecimensMutation is the GraphQL mutation that adds BloodSpecimen type objects to the Dgraph database.
	AddBloodSpecimensMutation = "mutation AddBloodSpecimens($input: [AddBloodSpecimenInput!]!) { addBloodSpecimen(input: $input) { bloodSpecimen { id } } }"

	// AddTestsMutation is the GraphQL mutation that adds Test type objects to the Dgraph database.
	AddTestsMutation = "mutation AddTests($input: [AddTestInput!]!) { addTest(input: $input) { test { id } } }"

	// AddResultsMutation is the GraphQL mutation that adds Result type objects to the Dgraph database.
	AddResultsMutation = "mutation AddResults($input: [AddResultInput!]!) { addResult(input: $input) { result { id } } }"
)
