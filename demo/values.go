package demo

var orgs = []string{
	"owner_org_a",
	"owner_org_b",
	"owner_org_c",
}

var labs = []string{
	"lab_org_a",
	"lab_org_b",
	"lab_org_c",
	"lab_org_d",
	"lab_org_e",
	"lab_org_f",
}

var storages = []string{
	"storage_org_a",
	"storage_org_b",
}

var streets = []string{
	"1st St.",
	"2nd St.",
	"3rd St.",
}

var cities = []string{
	"Bellefonte",
}

var counties = []string{
	"Centre",
}

var states = []string{
	"PA",
}

var zips = []int{
	16823,
}

var titles = []string{
	"Example Title A",
	"Example Title B",
}

var bodies = []string{
	"Example body content A",
	"Example body content B",
}

var protocolNames = []string{
	"Protocol A",
	"Protocol B",
	"Protocol C",
}

var planNames = []string{
	"Plan A",
	"Plan B",
	"Plan C",
}

var descriptions = []string{
	"It is a period of civil wars in the galaxy. A brave alliance of underground freedom fighters has challenged the tyranny and oppression of the awesome GALACTIC EMPIRE.",
	"Striking from a fortress hidden among the billion stars of the galaxy, rebel spaceships have won their first victory in a battle with the powerful Imperial Starfleet.",
	"The EMPIRE fears that another defeat could bring a thousand more solar systems into the rebellion, and Imperial control over the galaxy would be lost forever.",
	"To crush the rebellion once and for all, the EMPIRE is constructing a sinister new battle station. Powerful enough to destroy an entire planet, its completion spells certain doom for the champions of freedom.",
}

var notes = []string{
	"The Phantom Menace",
	"Attack of the Clones",
	"Revenge of the Sith",
	"A New Hope",
	"The Empire Strikes Back",
	"Return of the Jedi",
}

var users = []struct {
	email string
	first string
	last  string
	role  string
}{
	{
		email: "yoda@jedi.com",
		first: "yoda",
		last:  "---",
		role:  "USER_ADMIN",
	},
	{
		email: "dooku@jedi.com",
		first: "---",
		last:  "dooku",
		role:  "USER_INTERNAL",
	},
	{
		email: "jinn@jedi.com",
		first: "qui-gon",
		last:  "jinn",
		role:  "USER_LAB",
	},
	{
		email: "kenobi@jedi.com",
		first: "obi-wan",
		last:  "kenobi",
		role:  "USER_STORAGE",
	},
	{
		email: "skywalker@jedi.com",
		first: "anakin",
		last:  "skywalker",
		role:  "USER_STORAGE",
	},
	{
		email: "tano@jedi.com",
		first: "ashoka",
		last:  "tano",
		role:  "USER_LAB",
	},

	{
		email: "plagueis@sith.com",
		first: "hego",
		last:  "damask",
		role:  "USER_ADMIN",
	},
	{
		email: "sidious@sith.com",
		first: "sheev",
		last:  "palpatine",
		role:  "USER_ADMIN",
	},
	{
		email: "maul@sith.com",
		first: "maul",
		last:  "---",
		role:  "USER_LAB",
	},
	{
		email: "tyranus@sith.com",
		first: "---",
		last:  "dooku",
		role:  "USER_STORAGE",
	},
	{
		email: "vader@sith.com",
		first: "anakin",
		last:  "skywalker",
		role:  "USER_STORAGE",
	},
}

var ages = []int{
	20,
	30,
	40,
	50,
	60,
	70,
}
