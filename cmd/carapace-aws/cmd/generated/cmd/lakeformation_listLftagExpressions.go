package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listLftagExpressionsCmd = &cobra.Command{
	Use:   "list-lftag-expressions",
	Short: "Returns the LF-Tag expressions in caller’s account filtered based on caller's permissions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listLftagExpressionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_listLftagExpressionsCmd).Standalone()

		lakeformation_listLftagExpressionsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_listLftagExpressionsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		lakeformation_listLftagExpressionsCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve this list.")
	})
	lakeformationCmd.AddCommand(lakeformation_listLftagExpressionsCmd)
}
