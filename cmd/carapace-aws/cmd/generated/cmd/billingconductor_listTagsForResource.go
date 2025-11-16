package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductor_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "A list the tags for a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductor_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(billingconductor_listTagsForResourceCmd).Standalone()

		billingconductor_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that identifies the resource to list the tags.")
		billingconductor_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	billingconductorCmd.AddCommand(billingconductor_listTagsForResourceCmd)
}
