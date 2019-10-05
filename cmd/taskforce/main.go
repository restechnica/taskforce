package main

import (
	"fmt"
	"github.com/restechnica/taskforce/internal/hcl"
)

func main() {
	var config = hcl.Parse("../assets/example2.hcl")
	fmt.Printf("%+v\n", config)
}
