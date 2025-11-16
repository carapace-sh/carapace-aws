package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_tagResourceCmd).Standalone()

		guardduty_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the GuardDuty resource to apply a tag to.")
		guardduty_tagResourceCmd.Flags().String("tags", "", "The tags to be added to a resource.")
		guardduty_tagResourceCmd.MarkFlagRequired("resource-arn")
		guardduty_tagResourceCmd.MarkFlagRequired("tags")
	})
	guarddutyCmd.AddCommand(guardduty_tagResourceCmd)
}
