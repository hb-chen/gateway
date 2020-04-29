package registry

type Service struct {
	Name     string            `json:"name"`
	Version  string            `json:"version"`
	Metadata map[string]string `json:"metadata"`
	Methods  []*Method         `json:"methods"`
	Nodes    []*Node           `json:"nodes"`
}

type Node struct {
	Id       string            `json:"id"`
	Address  string            `json:"address"`
	Metadata map[string]string `json:"metadata"`
}

type Method struct {
	Name   string   `json:"name"`
	Routes []*Route `json:"routes"`
}

type Route struct {
	Method  string   `json:"method"`
	Pattern *Pattern `json:"pattern"`
}

type Pattern struct {
	Version         int      `json:"version"`
	Ops             []int    `json:"ops"`
	Pool            []string `json:"pool"`
	Verb            string   `json:"verb"`
	AssumeColonVerb bool     `json:"assume_colon_verb"`
}
