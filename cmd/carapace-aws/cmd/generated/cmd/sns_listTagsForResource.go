package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags added to the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_listTagsForResourceCmd).Standalone()

		sns_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the topic for which to list tags.")
		sns_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	snsCmd.AddCommand(sns_listTagsForResourceCmd)
}
