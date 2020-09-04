package graphql

// Sex is the enum defined for Donor fields in the schema.
var Sex = []string{
	"MALE",
	"FEMALE",
}

// Race is the enum defined for Donor fields in the schema.
var Race = []string{
	"AMERICAN_INDIAN_OR_ALASKA_NATIVE",
	"ASIAN",
	"BLACK_OR_AFRICAN_AMERICAN",
	"HISPANIC_OR_LATINO",
	"WHITE",
}

// SpecimenType is the enum defined for Specimen fields in the schema.
var SpecimenType = []string{
	"BLOOD",
}

// Container is the enum defined for Specimen fields in the schema.
var Container = []string{
	"VIAL",
}

// Status is the enum defined for Specimen fields in the schema.
var Status = []string{
	"DESTROYED",
	"EXHAUSTED",
	"IN_INVENTORY",
	"IN_TRANSIT",
	"LOST",
	"RESERVED",
	"TRANSFERRED",
}

// BloodType is the enum defined for BloodSpecimen fields in the schema.
var BloodType = []string{
	"O_NEG",
	"O_POS",
	"A_NEG",
	"A_POS",
	"B_NEG",
	"B_POS",
	"AB_NEG",
	"AB_POS",
}
