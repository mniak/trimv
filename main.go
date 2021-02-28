package main

import (
	"github.com/mniak/trimv/cmd"
	"github.com/mniak/trimv/internal"
)

func init() {
	cmd.RootCmd.Flags().StringP("input", "i", "", "Input file path")
	cmd.RootCmd.MarkFlagRequired("input")

	cmd.RootCmd.Flags().StringP("output", "o", "", "Output file path")
	cmd.RootCmd.MarkFlagRequired("output")

	cmd.RootCmd.Flags().Float32("intro-duration", 15, "Specifies the intro duration")

	cmd.RootCmd.Flags().Float32("outro-duration", 15, "Specifies the outro duration")
}

func main() {

	err := cmd.RootCmd.Execute()
	internal.Handle(err)
}
