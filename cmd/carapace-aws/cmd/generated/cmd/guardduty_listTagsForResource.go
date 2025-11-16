package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listTagsForResourceCmd).Standalone()

	guardduty_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the given GuardDuty resource.")
	guardduty_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	guarddutyCmd.AddCommand(guardduty_listTagsForResourceCmd)
}
