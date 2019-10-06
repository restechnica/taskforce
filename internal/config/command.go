package config

type Command struct {
	Name string `hcl:"name,label"`

	Expression string `hcl:"expression,attr"`
}

func (command Command) HasName(name string) bool {
	return command.Name == name
}
