package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "Amazon DynamoDB",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbCmd).Standalone()

	})
	rootCmd.AddCommand(dynamodbCmd)
}
