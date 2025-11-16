package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_listTagsOfResourceCmd = &cobra.Command{
	Use:   "list-tags-of-resource",
	Short: "List all tags on an Amazon DynamoDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_listTagsOfResourceCmd).Standalone()

	dynamodb_listTagsOfResourceCmd.Flags().String("next-token", "", "An optional string that, if supplied, must be copied from the output of a previous call to ListTagOfResource.")
	dynamodb_listTagsOfResourceCmd.Flags().String("resource-arn", "", "The Amazon DynamoDB resource with tags to be listed.")
	dynamodb_listTagsOfResourceCmd.MarkFlagRequired("resource-arn")
	dynamodbCmd.AddCommand(dynamodb_listTagsOfResourceCmd)
}
