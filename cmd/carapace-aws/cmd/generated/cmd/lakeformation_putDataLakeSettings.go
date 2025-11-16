package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lakeformation_putDataLakeSettingsCmd = &cobra.Command{
	Use:   "put-data-lake-settings",
	Short: "Sets the list of data lake administrators who have admin privileges on all resources managed by Lake Formation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lakeformation_putDataLakeSettingsCmd).Standalone()

	lakeformation_putDataLakeSettingsCmd.Flags().String("catalog-id", "", "The identifier for the Data Catalog.")
	lakeformation_putDataLakeSettingsCmd.Flags().String("data-lake-settings", "", "A structure representing a list of Lake Formation principals designated as data lake administrators.")
	lakeformation_putDataLakeSettingsCmd.MarkFlagRequired("data-lake-settings")
	lakeformationCmd.AddCommand(lakeformation_putDataLakeSettingsCmd)
}
