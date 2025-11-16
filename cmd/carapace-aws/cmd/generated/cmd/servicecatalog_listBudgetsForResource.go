package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_listBudgetsForResourceCmd = &cobra.Command{
	Use:   "list-budgets-for-resource",
	Short: "Lists all the budgets associated to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_listBudgetsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_listBudgetsForResourceCmd).Standalone()

		servicecatalog_listBudgetsForResourceCmd.Flags().String("accept-language", "", "The language code.")
		servicecatalog_listBudgetsForResourceCmd.Flags().String("page-size", "", "The maximum number of items to return with this call.")
		servicecatalog_listBudgetsForResourceCmd.Flags().String("page-token", "", "The page token for the next set of results.")
		servicecatalog_listBudgetsForResourceCmd.Flags().String("resource-id", "", "The resource identifier.")
		servicecatalog_listBudgetsForResourceCmd.MarkFlagRequired("resource-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_listBudgetsForResourceCmd)
}
