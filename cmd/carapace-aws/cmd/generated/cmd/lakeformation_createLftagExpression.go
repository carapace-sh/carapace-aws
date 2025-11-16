package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_createLftagExpressionCmd = &cobra.Command{
	Use:   "create-lftag-expression",
	Short: "Creates a new LF-Tag expression with the provided name, description, catalog ID, and expression body.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_createLftagExpressionCmd).Standalone()

	lakeformation_createLftagExpressionCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_createLftagExpressionCmd.Flags().String("description", "", "A description with information about the LF-Tag expression.")
	lakeformation_createLftagExpressionCmd.Flags().String("expression", "", "A list of LF-Tag conditions (key-value pairs).")
	lakeformation_createLftagExpressionCmd.Flags().String("name", "", "A name for the expression.")
	lakeformation_createLftagExpressionCmd.MarkFlagRequired("expression")
	lakeformation_createLftagExpressionCmd.MarkFlagRequired("name")
	lakeformationCmd.AddCommand(lakeformation_createLftagExpressionCmd)
}
