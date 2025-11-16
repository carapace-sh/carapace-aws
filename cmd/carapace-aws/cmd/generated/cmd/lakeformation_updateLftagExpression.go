package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_updateLftagExpressionCmd = &cobra.Command{
	Use:   "update-lftag-expression",
	Short: "Updates the name of the LF-Tag expression to the new description and expression body provided.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_updateLftagExpressionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_updateLftagExpressionCmd).Standalone()

		lakeformation_updateLftagExpressionCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_updateLftagExpressionCmd.Flags().String("description", "", "The description with information about the saved LF-Tag expression.")
		lakeformation_updateLftagExpressionCmd.Flags().String("expression", "", "The LF-Tag expression body composed of one more LF-Tag key-value pairs.")
		lakeformation_updateLftagExpressionCmd.Flags().String("name", "", "The name for the LF-Tag expression.")
		lakeformation_updateLftagExpressionCmd.MarkFlagRequired("expression")
		lakeformation_updateLftagExpressionCmd.MarkFlagRequired("name")
	})
	lakeformationCmd.AddCommand(lakeformation_updateLftagExpressionCmd)
}
