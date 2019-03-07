package storage

type App struct {
	Notes string `json:"notes"`
	UseFullEnvironment bool `json:"useFullEnvironment"`
	DetectAsRoot bool `json:"updateAsRoot"`
	DetectScript []string `json:"detectScript"`
	// DetectMatch string `json:"detectMatch"`
	PrepAsRoot bool `json:"updateAsRoot"`
	PrepScript []string `json:"prepScript"`
	UpdateAsRoot bool `json:"updateAsRoot"`
	UpdateScript []string `json:"updateScript"`
	PostAsRoot bool `json:"updateAsRoot"`
	PostScript []string `json:"updateScript"`
}

type Conf struct {
	RootCommand string `json:"rootCommand"`
	Apps map[string]App `json:"sites"`
}
