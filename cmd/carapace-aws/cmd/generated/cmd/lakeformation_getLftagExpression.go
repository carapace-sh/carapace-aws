package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getLftagExpressionCmd = &cobra.Command{
	Use:   "get-lftag-expression",
	Short: "Returns the details about the LF-Tag expression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getLftagExpressionCmd).Standalone()

	lakeformation_getLftagExpressionCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_getLftagExpressionCmd.Flags().String("name", "", "The name for the LF-Tag expression")
	lakeformation_getLftagExpressionCmd.MarkFlagRequired("name")
	lakeformationCmd.AddCommand(lakeformation_getLftagExpressionCmd)
}
