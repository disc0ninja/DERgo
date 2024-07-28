package internal

type Records struct {
	Records []Record `yaml:"records,flow"`
}

type Record struct {
	Name         string   `yaml:"name"`
	Expect       string   `yaml:"expect,omitempty"`
	Environments []string `yaml:"environments,flow,omitempty"`
}
