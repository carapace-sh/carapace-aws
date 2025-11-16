package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ssmQuicksetup_untagResourceCmd).Standalone()

		ssmQuicksetup_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to remove tags from.")
		ssmQuicksetup_untagResourceCmd.Flags().String("tag-keys", "", "The keys of the tags to remove from the resource.")
		ssmQuicksetup_untagResourceCmd.MarkFlagRequired("resource-arn")
		ssmQuicksetup_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_untagResourceCmd)
}
