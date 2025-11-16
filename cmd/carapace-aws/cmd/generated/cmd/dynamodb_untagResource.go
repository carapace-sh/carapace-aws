package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes the association of tags from an Amazon DynamoDB resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_untagResourceCmd).Standalone()

		dynamodb_untagResourceCmd.Flags().String("resource-arn", "", "The DynamoDB resource that the tags will be removed from.")
		dynamodb_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys.")
		dynamodb_untagResourceCmd.MarkFlagRequired("resource-arn")
		dynamodb_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	dynamodbCmd.AddCommand(dynamodb_untagResourceCmd)
}
