package thingfulx

// Category is a simple type alias for a string

const (

	// Permissions
	Reproduction = "https://creativecommons.org/ns#Reproduction"
	//making multiple copies

	Distribution = "https://creativecommons.org/ns#Distribution"
	//distribution, public display, and publicly performance

	DerivativeWorks = "https://creativecommons.org/ns#DerivativeWorks"
	//distribution of derivative works

	Sharing = "https://creativecommons.org/ns#Sharing"
	// permits commercial derivatives, but only non-commercial distribution

	// Requirements
	Notice = "https://creativecommons.org/ns#Notice"
	//copyright and license notices be kept intact

	Attribution = "https://creativecommons.org/ns#Attribution"
	// credit be given to copyright holder and/or author

	ShareAlike = "https://creativecommons.org/ns#ShareAlike"
	//derivative works be licensed under the same terms or compatible terms as the original work

	SourceCode = "https://creativecommons.org/ns#SourceCode"
	// source code (the preferred form for making modifications) must be provided when exercising some rights granted by the license.

	Copyleft = "https://creativecommons.org/ns#Copyleft"
	//derivative and combined works must be licensed under specified terms, similar to those on the original work

	LesserCopyleft = "https://creativecommons.org/ns#LesserCopyleft"
	//derivative works must be licensed under specified terms, with at least the same conditions as the original work; combinations with the work may be licensed under different terms

	// Prohibitions
	CommercialUse = "http://web.resource.org/cc/CommercialUse"
	// exercising rights for commercial purposes

	HighIncomeNationUse = "https://creativecommons.org/ns#HighIncomeNationUse"
	//use in a non-developing country

)

type License struct {
	// Human readable name for this license
	Name string `json:"name"`

	// url to the short version of the license
	URL string `json:"id"`

	// url to the long or legal version of the license
	LegalCodeURL string `json:"legalCodeUrl"`

	Permits []string `json:"permits"`

	Requires []string `json:"requires"`

	Prohibits []string `json:"prohibits"`
}

var (
	CC_0_V1 = License{

		Name: "CC0 1.0 Universal (CC0 1.0) Public Domain Dedication",

		URL: "https://creativecommons.org/publicdomain/zero/1.0/",

		LegalCodeURL: "https://creativecommons.org/publicdomain/zero/1.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{},

		Prohibits: []string{},
	}

	CC_BY_SA_V4 = License{

		Name: "Attribution-ShareAlike 4.0 International (CC BY-SA 4.0)",

		URL: "https://creativecommons.org/licenses/by-sa/4.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-sa/4.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
			ShareAlike,
		},

		Prohibits: []string{},
	}

	CC_BY_NC_V3 = License{

		Name: "Attribution-NonCommercial 3.0 Unported (CC BY-NC 3.0)",

		URL: "https://creativecommons.org/licenses/by-nc/3.0/",

		LegalCodeURL: "https://creativecommons.org/licenses/by-nc/3.0/legalcode",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{
			CommercialUse,
		},
	}

	OGL_V3 = License{

		Name: "open government licence version 3.0",

		URL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",

		LegalCodeURL: "http://www.nationalarchives.gov.uk/doc/open-government-licence/version/3/",

		Permits: []string{
			Reproduction,
			Distribution,
			DerivativeWorks,
			Sharing,
		},

		Requires: []string{
			Attribution,
		},

		Prohibits: []string{},
	}
)
