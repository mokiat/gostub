package main

import (
	"fmt"

	"github.com/codegangsta/cli"
)

type Stubber struct {
}

func (s *Stubber) Execute(c *cli.Context) {
	fmt.Println("Running `gostub`...")
}
