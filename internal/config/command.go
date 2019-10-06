package config

type Command struct {
	Name string `hcl:"name,label"`

	Expression string `hcl:"expression,attr"`
}
