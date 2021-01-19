package main

func init() {
	rootCmd.Flags().StringP("input", "i", "", "Input file path")
	rootCmd.MarkFlagRequired("input")

	rootCmd.Flags().StringP("output", "o", "", "Output file path")
	rootCmd.MarkFlagRequired("output")

	rootCmd.Flags().Float32("intro-duration", 15, "Specifies the intro duration")

	rootCmd.Flags().Float32("outro-duration", 15, "Specifies the outro duration")
}

func main() {

	err := rootCmd.Execute()
	handle(err)
}
