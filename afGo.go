package main

import (
	"afGo/afGoNet"
)

func main() {

	s := afGoNet.NewServer("afGo[1.0.2]")
	s.Server()
}
