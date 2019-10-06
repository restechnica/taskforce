package config

type Command struct {
	Name string `hcl:"name,label"`

	Executable string   `hcl:"executable,attr"`
	Arguments  []string `hcl:"arguments,attr"`
}
