package config

type Script struct {
	IsConcurrent bool          `hcl:"concurrent,optional"`
	Instructions []Instruction `hcl:"run,block"`
}
