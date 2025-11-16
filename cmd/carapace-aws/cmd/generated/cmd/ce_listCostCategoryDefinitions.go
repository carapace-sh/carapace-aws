package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_listCostCategoryDefinitionsCmd = &cobra.Command{
	Use:   "list-cost-category-definitions",
	Short: "Returns the name, Amazon Resource Name (ARN), `NumberOfRules` and effective dates of all Cost Categories defined in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_listCostCategoryDefinitionsCmd).Standalone()

	ce_listCostCategoryDefinitionsCmd.Flags().String("effective-on", "", "The date when the Cost Category was effective.")
	ce_listCostCategoryDefinitionsCmd.Flags().String("max-results", "", "The number of entries a paginated response contains.")
	ce_listCostCategoryDefinitionsCmd.Flags().String("next-token", "", "The token to retrieve the next set of results.")
	ceCmd.AddCommand(ce_listCostCategoryDefinitionsCmd)
}
