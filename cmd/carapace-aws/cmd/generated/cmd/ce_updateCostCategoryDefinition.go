package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_updateCostCategoryDefinitionCmd = &cobra.Command{
	Use:   "update-cost-category-definition",
	Short: "Updates an existing Cost Category.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_updateCostCategoryDefinitionCmd).Standalone()

	ce_updateCostCategoryDefinitionCmd.Flags().String("cost-category-arn", "", "The unique identifier for your Cost Category.")
	ce_updateCostCategoryDefinitionCmd.Flags().String("default-value", "", "")
	ce_updateCostCategoryDefinitionCmd.Flags().String("effective-start", "", "The Cost Category's effective start date.")
	ce_updateCostCategoryDefinitionCmd.Flags().String("rule-version", "", "")
	ce_updateCostCategoryDefinitionCmd.Flags().String("rules", "", "The `Expression` object used to categorize costs.")
	ce_updateCostCategoryDefinitionCmd.Flags().String("split-charge-rules", "", "The split charge rules used to allocate your charges between your Cost Category values.")
	ce_updateCostCategoryDefinitionCmd.MarkFlagRequired("cost-category-arn")
	ce_updateCostCategoryDefinitionCmd.MarkFlagRequired("rule-version")
	ce_updateCostCategoryDefinitionCmd.MarkFlagRequired("rules")
	ceCmd.AddCommand(ce_updateCostCategoryDefinitionCmd)
}
