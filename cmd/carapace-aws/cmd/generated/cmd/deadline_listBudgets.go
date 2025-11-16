package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deadline_listBudgetsCmd = &cobra.Command{
	Use:   "list-budgets",
	Short: "A list of budgets in a farm.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deadline_listBudgetsCmd).Standalone()

	deadline_listBudgetsCmd.Flags().String("farm-id", "", "The farm ID associated with the budgets.")
	deadline_listBudgetsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
	deadline_listBudgetsCmd.Flags().String("next-token", "", "The token for the next set of results, or `null` to start from the beginning.")
	deadline_listBudgetsCmd.Flags().String("status", "", "The status to list for the budgets.")
	deadline_listBudgetsCmd.MarkFlagRequired("farm-id")
	deadlineCmd.AddCommand(deadline_listBudgetsCmd)
}
