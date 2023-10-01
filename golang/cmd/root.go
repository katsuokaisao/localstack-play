package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use: "",
}

func init() {
	rootCmd.AddCommand(
		s3playCmd,
	)
}

func Execute() error {
	return rootCmd.Execute()
}
