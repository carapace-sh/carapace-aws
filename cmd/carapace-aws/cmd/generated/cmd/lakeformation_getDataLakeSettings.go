package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_getDataLakeSettingsCmd = &cobra.Command{
	Use:   "get-data-lake-settings",
	Short: "Retrieves the list of the data lake administrators of a Lake Formation-managed data lake.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_getDataLakeSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lakeformation_getDataLakeSettingsCmd).Standalone()

		lakeformation_getDataLakeSettingsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	})
	lakeformationCmd.AddCommand(lakeformation_getDataLakeSettingsCmd)
}
