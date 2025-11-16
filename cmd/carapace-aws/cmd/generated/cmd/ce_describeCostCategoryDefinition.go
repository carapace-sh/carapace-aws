package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_describeCostCategoryDefinitionCmd = &cobra.Command{
	Use:   "describe-cost-category-definition",
	Short: "Returns the name, Amazon Resource Name (ARN), rules, definition, and effective dates of a Cost Category that's defined in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_describeCostCategoryDefinitionCmd).Standalone()

	ce_describeCostCategoryDefinitionCmd.Flags().String("cost-category-arn", "", "The unique identifier for your Cost Category.")
	ce_describeCostCategoryDefinitionCmd.Flags().String("effective-on", "", "The date when the Cost Category was effective.")
	ce_describeCostCategoryDefinitionCmd.MarkFlagRequired("cost-category-arn")
	ceCmd.AddCommand(ce_describeCostCategoryDefinitionCmd)
}
