package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes tags associated with a budget or budget action resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_untagResourceCmd).Standalone()

		budgets_untagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
		budgets_untagResourceCmd.Flags().String("resource-tag-keys", "", "The key that's associated with the tag.")
		budgets_untagResourceCmd.MarkFlagRequired("resource-arn")
		budgets_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	})
	budgetsCmd.AddCommand(budgets_untagResourceCmd)
}
