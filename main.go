package main

import (
	"github.com/mniak/trimv/cmd"
	"github.com/mniak/trimv/internal"
)

func main() {
	err := cmd.RootCmd.Execute()
	internal.Handle(err)
}
