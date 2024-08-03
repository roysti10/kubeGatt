package structs

type Pipeline struct {
	Stages map[string][]string `yaml:"stages"`
	Steps  map[string]Step     `yaml:"steps"`
}

type Step struct {
	Image   string `yaml:"image"`
	Command string `yaml:"command"`
}
