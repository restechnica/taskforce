package config

type Task struct {
	Name         string        `hcl:"name,label"`
	Instructions []Instruction `hcl:"run,block"`
}

func (task Task) HasName(name string) bool {
	return task.Name == name
}
