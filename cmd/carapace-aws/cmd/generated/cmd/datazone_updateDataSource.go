package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_updateDataSourceCmd = &cobra.Command{
	Use:   "update-data-source",
	Short: "Updates the specified data source in Amazon DataZone.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_updateDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_updateDataSourceCmd).Standalone()

		datazone_updateDataSourceCmd.Flags().String("asset-forms-input", "", "The asset forms to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().String("configuration", "", "The configuration to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().String("description", "", "The description to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().String("domain-identifier", "", "The identifier of the domain in which to update a data source.")
		datazone_updateDataSourceCmd.Flags().String("enable-setting", "", "The enable setting to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().String("identifier", "", "The identifier of the data source to be updated.")
		datazone_updateDataSourceCmd.Flags().String("name", "", "The name to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().Bool("no-publish-on-import", false, "The publish on import setting to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().Bool("no-retain-permissions-on-revoke-failure", false, "Specifies that the granted permissions are retained in case of a self-subscribe functionality failure for a data source.")
		datazone_updateDataSourceCmd.Flags().Bool("publish-on-import", false, "The publish on import setting to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().String("recommendation", "", "The recommendation to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.Flags().Bool("retain-permissions-on-revoke-failure", false, "Specifies that the granted permissions are retained in case of a self-subscribe functionality failure for a data source.")
		datazone_updateDataSourceCmd.Flags().String("schedule", "", "The schedule to be updated as part of the `UpdateDataSource` action.")
		datazone_updateDataSourceCmd.MarkFlagRequired("domain-identifier")
		datazone_updateDataSourceCmd.MarkFlagRequired("identifier")
		datazone_updateDataSourceCmd.Flag("no-publish-on-import").Hidden = true
		datazone_updateDataSourceCmd.Flag("no-retain-permissions-on-revoke-failure").Hidden = true
	})
	datazoneCmd.AddCommand(datazone_updateDataSourceCmd)
}
