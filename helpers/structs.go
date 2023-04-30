package helpers

type ItemInfo struct {
	Item string ``
	Info string ``
}

type CustomizationStruct struct {
	Themes []Themes `json:"Themes,omitempty"`
}

type Themes struct {
	Name     string                 `json:"Name"`
	Location string                 `json:"Location"`
	Packages map[string]interface{} `json:"Packages,omitempty"`
	Assets   map[string]string      `json:"Assets,omitempty"`
}

type MyDictionary struct {
	Data map[string]map[string]string
}

type Instruction struct {
	Command          string   `json:"Command"`
	UseSudo          bool     `json:"UseSudo"`
	From             string   `json:"from,omitempty"`
	To               string   `json:"to,omitempty"`
	Source           string   `json:"Source,omitempty"`
	NewName          string   `json:"NewName,omitempty"`
	SourceLine       string   `json:"SourceLine,omitempty"`
	EditedLine       string   `json:"EditedLine,omitempty"`
	FileReplace      string   `json:"FileReplace,omitempty"`
	GitURL           string   `json:"GitURL,omitempty"`
	GitDestination   string   `json:"GitDestination,omitempty"`
	CommandDirectory string   `json:"CommandDirectory,omitempty"`
	Args             []string `json:"Args,omitempty"`
}

type Instructions struct {
	Instructions []Instruction `json:"Instructions"`
}
