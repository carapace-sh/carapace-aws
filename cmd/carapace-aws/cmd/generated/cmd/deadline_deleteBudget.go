package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_deleteBudgetCmd = &cobra.Command{
	Use:   "delete-budget",
	Short: "Deletes a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_deleteBudgetCmd).Standalone()

	deadline_deleteBudgetCmd.Flags().String("budget-id", "", "The budget ID of the budget to delete.")
	deadline_deleteBudgetCmd.Flags().String("farm-id", "", "The farm ID of the farm to remove from the budget.")
	deadline_deleteBudgetCmd.MarkFlagRequired("budget-id")
	deadline_deleteBudgetCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_deleteBudgetCmd)
}
