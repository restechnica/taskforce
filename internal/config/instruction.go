package config

type Instruction struct {
	Type      string `hcl:"type,label"`
	Reference string `hcl:"reference,label"`
}
