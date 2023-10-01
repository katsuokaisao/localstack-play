package cmd

import (
	"github.com/katsuokaisao/localstack/inject/s3play"
	"github.com/spf13/cobra"
)

var s3playCmd = &cobra.Command{
	Use:   "s3play",
	Short: "s3play is a command line tool for playing with S3",
	Long: `s3play is a command line tool for playing with S3.
It is a simple wrapper around the AWS SDK for Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		app, err := s3play.SetApplication()
		if err != nil {
			panic(err)
		}

		app.Run()
	},
}

