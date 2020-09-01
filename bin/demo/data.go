package main

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

var countries = []string{
	"United States",
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
	"Turmoil has engulfed the Galactic Republic. The taxation of trade routes to outlying star systems is in dispute.",
	"Hoping to resolve the matter with a blockade of deadly battleships, the greedy Trade Federation has stopped all shipping to the small planet of Naboo.",
	"While the congress of the Republic endlessly debates this alarming chain of events, the Supreme Chancellor has secretly dispatched two Jedi Knights, the guardians of peace and justice in the galaxy, to settle the conflict....",
	"There is unrest in the Galactic Senate. Several thousand solar systems have declared their intentions to leave the Republic.",
	"This Separatist movement, under the leadership of the mysterious Count Dooku, has made it difficult for the limited number of Jedi Knights to maintain peace and order in the galaxy.",
	"Senator Amidala, the former Queen of Naboo, is returning to the Galactic Senate to vote on the critical issue of creating an ARMY OF THE REPUBLIC to assist the overwhelmed Jedi....",
	"War! The Republic is crumbling under attacks by the ruthless Sith Lord, Count Dooku. There are heroes on both sides. Evil is everywhere.",
	"In a stunning move, the fiendish droid leader, General Grievous, has swept into the Republic capital and kidnapped Chancellor Palpatine, leader of the Galactic Senate.",
	"As the Separatist Droid Army attempts to flee the besieged capital with their valuable hostage, two Jedi Knights lead a desperate mission to rescue the captive Chancellor....",
	"It is a period of civil war. Rebel spaceships, striking from a hidden base, have won their first victory against the evil Galactic Empire.",
	"During the battle, Rebel spies managed to steal secret plans to the Empire's ultimate weapon, the DEATH STAR, an armored space station with enough power to destroy an entire planet.",
	"Pursued by the Empire's sinister agents, Princess Leia races home aboard her starship, custodian of the stolen plans that can save her people and restore freedom to the galaxy....",
	"It is a dark time for the Rebellion. Although the Death Star has been destroyed, Imperial troops have driven the Rebel forces from their hidden base and pursued them across the galaxy.",
	"Evading the dreaded Imperial Starfleet, a group of freedom fighters led by Luke Skywalker have established a new secret base on the remote ice world of Hoth.",
	"The evil lord Darth Vader, obsessed with finding young Skywalker, has dispatched thousands of remote probes into the far reaches of space....",
	"Luke Skywalker has returned to his home planet of Tatooine in an attempt to rescue his friend Han Solo from the clutches of the vile gangster Jabba the Hutt.",
	"Little does Luke know that the GALACTIC EMPIRE has secretly begun construction on a new armored space station even more powerful than the first dreaded Death Star.",
	"When completed, this ultimate weapon will spell certain doom for the small band of rebels struggling to restore freedom to the galaxy...",
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
	email  string
	first  string
	last   string
	role   string
	userID string
}{
	{
		email:  "yoda@jedi.com",
		first:  "yoda",
		last:   "---",
		role:   "USER_ADMIN",
		userID: "auth0|123",
	},
	{
		email:  "dooku@jedi.com",
		first:  "---",
		last:   "dooku",
		role:   "USER_INTERNAL",
		userID: "auth0|234",
	},
	{
		email:  "jinn@jedi.com",
		first:  "qui-gon",
		last:   "jinn",
		role:   "USER_LAB",
		userID: "auth0|345",
	},
	{
		email:  "kenobi@jedi.com",
		first:  "obi-wan",
		last:   "kenobi",
		role:   "USER_STORAGE",
		userID: "auth0|456",
	},
	{
		email:  "skywalker@jedi.com",
		first:  "anakin",
		last:   "skywalker",
		role:   "USER_STORAGE",
		userID: "auth0|567",
	},
	{
		email:  "tano@jedi.com",
		first:  "ashoka",
		last:   "tano",
		role:   "USER_LAB",
		userID: "auth0|678",
	},
	{
		email:  "plagueis@sith.com",
		first:  "hego",
		last:   "damask",
		role:   "USER_ADMIN",
		userID: "auth0|789",
	},
	{
		email:  "sidious@sith.com",
		first:  "sheev",
		last:   "palpatine",
		role:   "USER_ADMIN",
		userID: "auth0|890",
	},
	{
		email:  "maul@sith.com",
		first:  "maul",
		last:   "---",
		role:   "USER_LAB",
		userID: "auth0|321",
	},
	{
		email:  "tyranus@sith.com",
		first:  "---",
		last:   "dooku",
		role:   "USER_STORAGE",
		userID: "auth0|432",
	},
	{
		email:  "vader@sith.com",
		first:  "anakin",
		last:   "skywalker",
		role:   "USER_STORAGE",
		userID: "auth0|543",
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
