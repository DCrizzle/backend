package loader

const addOwnerOrgsMutation = "mutation AddOwnerOrgs($input: [AddOwnerOrgInput!]!) { addOwnerOrg(input: $input) { ownerOrg { id } } }"

const addLabOrgsMutation = "mutation AddLabOrgs($input: [AddLabOrgInput!]!) { addLabOrg(input: $input) { labOrg { id } } }"

const addStorageOrgsMutation = "mutation addStorageOrg(input: $input) { id }"

const addUsersMutation = "mutation addUser(input: $input) { id }"

const addProtocolsMutation = "mutation addProtocol(input: $input) { id }"

const addProtocolFormsMutation = "mutation addProtocolForm(input: $input) { id }"

const addPlansMutation = "mutation addPlans(input: $input) { id }"

const addConsentFormsMutation = "mutation addConsentForm(input: $input) { id }"

const addDonorsMutation = "mutation addDonor(input: $input) { id }"

const addConsentsMutation = "mutation addConsent(input: $input) { id }"

const addBloodSpecimensMutation = "mutation addBloodSpecimen(input: $input) { id }"

const addTestsMutation = "mutation addTest(input: $input) { id }"

const addResultsMutation = "mutation addResult(input: $input) { id }"
