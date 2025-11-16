package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sqsCmd = &cobra.Command{
	Use:   "sqs",
	Short: "Welcome to the *Amazon SQS API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sqsCmd).Standalone()

	rootCmd.AddCommand(sqsCmd)
}
