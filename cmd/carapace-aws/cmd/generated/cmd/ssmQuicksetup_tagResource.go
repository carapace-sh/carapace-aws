package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns key-value pairs of metadata to Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_tagResourceCmd).Standalone()

		ssmQuicksetup_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to tag.")
		ssmQuicksetup_tagResourceCmd.Flags().String("tags", "", "Key-value pairs of metadata to assign to the resource.")
		ssmQuicksetup_tagResourceCmd.MarkFlagRequired("resource-arn")
		ssmQuicksetup_tagResourceCmd.MarkFlagRequired("tags")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_tagResourceCmd)
}
