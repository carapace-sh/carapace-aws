package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_deleteCostCategoryDefinitionCmd = &cobra.Command{
	Use:   "delete-cost-category-definition",
	Short: "Deletes a Cost Category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_deleteCostCategoryDefinitionCmd).Standalone()

	ce_deleteCostCategoryDefinitionCmd.Flags().String("cost-category-arn", "", "The unique identifier for your Cost Category.")
	ce_deleteCostCategoryDefinitionCmd.MarkFlagRequired("cost-category-arn")
	ceCmd.AddCommand(ce_deleteCostCategoryDefinitionCmd)
}
