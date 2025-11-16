package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicecatalog_disassociateBudgetFromResourceCmd = &cobra.Command{
	Use:   "disassociate-budget-from-resource",
	Short: "Disassociates the specified budget from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicecatalog_disassociateBudgetFromResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicecatalog_disassociateBudgetFromResourceCmd).Standalone()

		servicecatalog_disassociateBudgetFromResourceCmd.Flags().String("budget-name", "", "The name of the budget you want to disassociate.")
		servicecatalog_disassociateBudgetFromResourceCmd.Flags().String("resource-id", "", "The resource identifier you want to disassociate from.")
		servicecatalog_disassociateBudgetFromResourceCmd.MarkFlagRequired("budget-name")
		servicecatalog_disassociateBudgetFromResourceCmd.MarkFlagRequired("resource-id")
	})
	servicecatalogCmd.AddCommand(servicecatalog_disassociateBudgetFromResourceCmd)
}
