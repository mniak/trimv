package main

const INTRO_SIZE = 15

func init() {
	rootCmd.Flags().StringP("input", "i", "", "Input file path")
	rootCmd.MarkFlagRequired("input")

	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.MarkFlagRequired("output")
}

func main() {

	err := rootCmd.Execute()
	handle(err)
}
