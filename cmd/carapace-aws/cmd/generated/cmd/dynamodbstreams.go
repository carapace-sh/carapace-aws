package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodbstreamsCmd = &cobra.Command{
	Use:   "dynamodbstreams",
	Short: "Amazon DynamoDB\n\nAmazon DynamoDB Streams provides API actions for accessing streams and processing stream records.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodbstreamsCmd).Standalone()

	rootCmd.AddCommand(dynamodbstreamsCmd)
}
