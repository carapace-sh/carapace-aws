package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getResourceLftagsCmd = &cobra.Command{
	Use:   "get-resource-lftags",
	Short: "Returns the LF-tags applied to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getResourceLftagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getResourceLftagsCmd).Standalone()

		lakeformation_getResourceLftagsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
		lakeformation_getResourceLftagsCmd.Flags().String("resource", "", "The database, table, or column resource for which you want to return LF-tags.")
		lakeformation_getResourceLftagsCmd.Flags().String("show-assigned-lftags", "", "Indicates whether to show the assigned LF-tags.")
		lakeformation_getResourceLftagsCmd.MarkFlagRequired("resource")
	})
	lakeformationCmd.AddCommand(lakeformation_getResourceLftagsCmd)
}
