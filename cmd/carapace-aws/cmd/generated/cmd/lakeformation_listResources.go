package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_listResourcesCmd = &cobra.Command{
	Use:   "list-resources",
	Short: "Lists the resources registered to be managed by the Data Catalog.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_listResourcesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_listResourcesCmd).Standalone()

		lakeformation_listResourcesCmd.Flags().String("filter-condition-list", "", "Any applicable row-level and/or column-level filtering conditions for the resources.")
		lakeformation_listResourcesCmd.Flags().String("max-results", "", "The maximum number of resource results.")
		lakeformation_listResourcesCmd.Flags().String("next-token", "", "A continuation token, if this is not the first call to retrieve these resources.")
	})
	lakeformationCmd.AddCommand(lakeformation_listResourcesCmd)
}
