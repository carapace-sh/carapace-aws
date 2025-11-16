package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_getBudgetCmd = &cobra.Command{
	Use:   "get-budget",
	Short: "Get a budget.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_getBudgetCmd).Standalone()

	deadline_getBudgetCmd.Flags().String("budget-id", "", "The budget ID.")
	deadline_getBudgetCmd.Flags().String("farm-id", "", "The farm ID of the farm connected to the budget.")
	deadline_getBudgetCmd.MarkFlagRequired("budget-id")
	deadline_getBudgetCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_getBudgetCmd)
}
