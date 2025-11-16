package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_associateBudgetWithResourceCmd = &cobra.Command{
	Use:   "associate-budget-with-resource",
	Short: "Associates the specified budget with the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_associateBudgetWithResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_associateBudgetWithResourceCmd).Standalone()

		servicecatalog_associateBudgetWithResourceCmd.Flags().String("budget-name", "", "The name of the budget you want to associate.")
		servicecatalog_associateBudgetWithResourceCmd.Flags().String("resource-id", "", "The resource identifier.")
		servicecatalog_associateBudgetWithResourceCmd.MarkFlagRequired("budget-name")
		servicecatalog_associateBudgetWithResourceCmd.MarkFlagRequired("resource-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_associateBudgetWithResourceCmd)
}
