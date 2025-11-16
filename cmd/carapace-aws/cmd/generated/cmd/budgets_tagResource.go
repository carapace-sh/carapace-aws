package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates tags for a budget or budget action resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(budgets_tagResourceCmd).Standalone()

		budgets_tagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
		budgets_tagResourceCmd.Flags().String("resource-tags", "", "The tags associated with the resource.")
		budgets_tagResourceCmd.MarkFlagRequired("resource-arn")
		budgets_tagResourceCmd.MarkFlagRequired("resource-tags")
	})
	budgetsCmd.AddCommand(budgets_tagResourceCmd)
}
