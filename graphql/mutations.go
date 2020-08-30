package graphql

const addOwnerOrgsMutation = "mutation AddOwnerOrgs($input: [AddOwnerOrgInput!]!) { addOwnerOrg(input: $input) { ownerOrg { id } } }"

const addLabOrgsMutation = "mutation AddLabOrgs($input: [AddLabOrgInput!]!) { addLabOrg(input: $input) { labOrg { id } } }"

const addStorageOrgsMutation = "mutation AddStorageOrgs($input: [AddStorageOrgInput!]!) { addStorageOrg(input: $input) { storageOrg { id } } }"

const addUsersMutation = "mutation AddUsers($input: [AddUserInput!]!) { addUser(input: $input) { user { email } } }"

const updateUserMutation = "mutation EditUser($input: UpdateUserInput!) { updateUser(input: $input) { user { email } } }"

const deleteUserMutation = "mutation RemoveUser($filter: UserFilter!) { deleteUser(filter: $filter) { user { email } } }"

const addProtocolsMutation = "mutation AddProtocols($input: [AddProtocolInput!]!) { addProtocol(input: $input) { protocol { id } } }"

const addProtocolFormsMutation = "mutation AddProtocolForms($input: [AddProtocolFormInput!]!) { addProtocolForm(input: $input) { protocolForm { id } } }"

const addPlansMutation = "mutation AddPlans($input: [AddPlanInput!]!) { addPlan(input: $input) { plan { id } } }"

const addConsentFormsMutation = "mutation AddConsentForms($input: [AddConsentFormInput!]!) { addConsentForm(input: $input) { consentForm { id } } }"

const addDonorsMutation = "mutation AddDonors($input: [AddDonorInput!]!) { addDonor(input: $input) { donor { id } } }"

const addConsentsMutation = "mutation AddConsent($input: [AddConsentInput!]!) { addConsent(input: $input) { consent { id } } }"

const addBloodSpecimensMutation = "mutation AddBloodSpecimens($input: [AddBloodSpecimenInput!]!) { addBloodSpecimen(input: $input) { bloodSpecimen { id } } }"

const addTestsMutation = "mutation AddTests($input: [AddTestInput!]!) { addTest(input: $input) { test { id } } }"

const addResultsMutation = "mutation AddResults($input: [AddResultInput!]!) { addResult(input: $input) { result { id } } }"
