package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sns_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add tags to the specified Amazon SNS topic.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sns_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sns_tagResourceCmd).Standalone()

		sns_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the topic to which to add tags.")
		sns_tagResourceCmd.Flags().String("tags", "", "The tags to be added to the specified topic.")
		sns_tagResourceCmd.MarkFlagRequired("resource-arn")
		sns_tagResourceCmd.MarkFlagRequired("tags")
	})
	snsCmd.AddCommand(sns_tagResourceCmd)
}
