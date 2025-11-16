package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associate a set of tags with an Amazon DynamoDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_tagResourceCmd).Standalone()

		dynamodb_tagResourceCmd.Flags().String("resource-arn", "", "Identifies the Amazon DynamoDB resource to which tags should be added.")
		dynamodb_tagResourceCmd.Flags().String("tags", "", "The tags to be assigned to the Amazon DynamoDB resource.")
		dynamodb_tagResourceCmd.MarkFlagRequired("resource-arn")
		dynamodb_tagResourceCmd.MarkFlagRequired("tags")
	})
	dynamodbCmd.AddCommand(dynamodb_tagResourceCmd)
}
