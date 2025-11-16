package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var budgets_describeBudgetsCmd = &cobra.Command{
	Use:   "describe-budgets",
	Short: "Lists the budgets that are associated with an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(budgets_describeBudgetsCmd).Standalone()

	budgets_describeBudgetsCmd.Flags().String("account-id", "", "The `accountId` that is associated with the budgets that you want to describe.")
	budgets_describeBudgetsCmd.Flags().String("max-results", "", "An integer that represents how many budgets a paginated response contains.")
	budgets_describeBudgetsCmd.Flags().String("next-token", "", "The pagination token that you include in your request to indicate the next set of results that you want to retrieve.")
	budgets_describeBudgetsCmd.Flags().String("show-filter-expression", "", "Specifies whether the response includes the filter expression associated with the budgets.")
	budgets_describeBudgetsCmd.MarkFlagRequired("account-id")
	budgetsCmd.AddCommand(budgets_describeBudgetsCmd)
}
