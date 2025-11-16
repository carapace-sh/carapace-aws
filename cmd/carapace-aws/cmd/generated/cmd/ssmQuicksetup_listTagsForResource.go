package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ssmQuicksetup_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns tags assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ssmQuicksetup_listTagsForResourceCmd).Standalone()

	ssmQuicksetup_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource the tag is assigned to.")
	ssmQuicksetup_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ssmQuicksetupCmd.AddCommand(ssmQuicksetup_listTagsForResourceCmd)
}
