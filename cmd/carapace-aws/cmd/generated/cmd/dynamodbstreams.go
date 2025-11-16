package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreamsCmd = &cobra.Command{
	Use:   "dynamodbstreams",
	Short: "Amazon DynamoDB",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreamsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbstreamsCmd).Standalone()

	})
	rootCmd.AddCommand(dynamodbstreamsCmd)
}
