package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_createCostCategoryDefinitionCmd = &cobra.Command{
	Use:   "create-cost-category-definition",
	Short: "Creates a new Cost Category with the requested name and rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_createCostCategoryDefinitionCmd).Standalone()

	ce_createCostCategoryDefinitionCmd.Flags().String("default-value", "", "")
	ce_createCostCategoryDefinitionCmd.Flags().String("effective-start", "", "The Cost Category's effective start date.")
	ce_createCostCategoryDefinitionCmd.Flags().String("name", "", "")
	ce_createCostCategoryDefinitionCmd.Flags().String("resource-tags", "", "An optional list of tags to associate with the specified [`CostCategory`](https://docs.aws.amazon.com/aws-cost-management/latest/APIReference/API_CostCategory.html) .")
	ce_createCostCategoryDefinitionCmd.Flags().String("rule-version", "", "")
	ce_createCostCategoryDefinitionCmd.Flags().String("rules", "", "The Cost Category rules used to categorize costs.")
	ce_createCostCategoryDefinitionCmd.Flags().String("split-charge-rules", "", "The split charge rules used to allocate your charges between your Cost Category values.")
	ce_createCostCategoryDefinitionCmd.MarkFlagRequired("name")
	ce_createCostCategoryDefinitionCmd.MarkFlagRequired("rule-version")
	ce_createCostCategoryDefinitionCmd.MarkFlagRequired("rules")
	ceCmd.AddCommand(ce_createCostCategoryDefinitionCmd)
}
