package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_createBudgetCmd = &cobra.Command{
	Use:   "create-budget",
	Short: "Creates a budget to set spending thresholds for your rendering activity.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_createBudgetCmd).Standalone()

	deadline_createBudgetCmd.Flags().String("actions", "", "The budget actions to specify what happens when the budget runs out.")
	deadline_createBudgetCmd.Flags().String("approximate-dollar-limit", "", "The dollar limit based on consumed usage.")
	deadline_createBudgetCmd.Flags().String("client-token", "", "The unique token which the server uses to recognize retries of the same request.")
	deadline_createBudgetCmd.Flags().String("description", "", "The description of the budget.")
	deadline_createBudgetCmd.Flags().String("display-name", "", "The display name of the budget.")
	deadline_createBudgetCmd.Flags().String("farm-id", "", "The farm ID to include in this budget.")
	deadline_createBudgetCmd.Flags().String("schedule", "", "The schedule to associate with this budget.")
	deadline_createBudgetCmd.Flags().String("usage-tracking-resource", "", "The queue ID provided to this budget to track usage.")
	deadline_createBudgetCmd.MarkFlagRequired("actions")
	deadline_createBudgetCmd.MarkFlagRequired("approximate-dollar-limit")
	deadline_createBudgetCmd.MarkFlagRequired("display-name")
	deadline_createBudgetCmd.MarkFlagRequired("farm-id")
	deadline_createBudgetCmd.MarkFlagRequired("schedule")
	deadline_createBudgetCmd.MarkFlagRequired("usage-tracking-resource")
	deadlineCmd.AddCommand(deadline_createBudgetCmd)
}
