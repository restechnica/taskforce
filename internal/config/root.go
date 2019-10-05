package config

type Root struct {
	Commands []Command `hcl:"command,block"`
}
