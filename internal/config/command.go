package config

type Command struct {
	Name      string `hcl:"name,label"`
	Directory string `hcl:"directory,optional"`
	Text      string `hcl:"text,attr"`
}

func (command Command) HasName(name string) bool {
	return command.Name == name
}

func (command Command) HasDirectory() bool {
	return command.Directory != ""
}
