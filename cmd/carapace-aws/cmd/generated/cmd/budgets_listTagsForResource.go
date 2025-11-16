package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags associated with a budget or budget action resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_listTagsForResourceCmd).Standalone()

	budgets_listTagsForResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
	budgets_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	budgetsCmd.AddCommand(budgets_listTagsForResourceCmd)
}
