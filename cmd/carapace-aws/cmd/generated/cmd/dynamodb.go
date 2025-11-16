package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbCmd = &cobra.Command{
	Use:   "dynamodb",
	Short: "Amazon DynamoDB\n\nAmazon DynamoDB is a fully managed NoSQL database service that provides fast and predictable performance with seamless scalability.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodbCmd).Standalone()

	})
	rootCmd.AddCommand(dynamodbCmd)
}
