package config

type Task struct {
	Name    string   `hcl:"name,label"`
	Scripts []Script `hcl:"script,block"`
}

func (task Task) HasName(name string) bool {
	return task.Name == name
}
