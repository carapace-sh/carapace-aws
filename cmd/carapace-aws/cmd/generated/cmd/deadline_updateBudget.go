package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_updateBudgetCmd = &cobra.Command{
	Use:   "update-budget",
	Short: "Updates a budget that sets spending thresholds for rendering activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_updateBudgetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(deadline_updateBudgetCmd).Standalone()

		deadline_updateBudgetCmd.Flags().String("actions-to-add", "", "The budget actions to add.")
		deadline_updateBudgetCmd.Flags().String("actions-to-remove", "", "The budget actions to remove from the budget.")
		deadline_updateBudgetCmd.Flags().String("approximate-dollar-limit", "", "The dollar limit to update on the budget.")
		deadline_updateBudgetCmd.Flags().String("budget-id", "", "The budget ID to update.")
		deadline_updateBudgetCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
		deadline_updateBudgetCmd.Flags().String("description", "", "The description of the budget to update.")
		deadline_updateBudgetCmd.Flags().String("display-name", "", "The display name of the budget to update.")
		deadline_updateBudgetCmd.Flags().String("farm-id", "", "The farm ID of the budget to update.")
		deadline_updateBudgetCmd.Flags().String("schedule", "", "The schedule to update.")
		deadline_updateBudgetCmd.Flags().String("status", "", "Updates the status of the budget.")
		deadline_updateBudgetCmd.MarkFlagRequired("budget-id")
		deadline_updateBudgetCmd.MarkFlagRequired("farm-id")
	})
	deadlineCmd.AddCommand(deadline_updateBudgetCmd)
}
