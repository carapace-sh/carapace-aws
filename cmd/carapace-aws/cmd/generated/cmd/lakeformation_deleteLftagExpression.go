package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_deleteLftagExpressionCmd = &cobra.Command{
	Use:   "delete-lftag-expression",
	Short: "Deletes the LF-Tag expression.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_deleteLftagExpressionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_deleteLftagExpressionCmd).Standalone()

		lakeformation_deleteLftagExpressionCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_deleteLftagExpressionCmd.Flags().String("name", "", "The name for the LF-Tag expression.")
		lakeformation_deleteLftagExpressionCmd.MarkFlagRequired("name")
	})
	lakeformationCmd.AddCommand(lakeformation_deleteLftagExpressionCmd)
}
