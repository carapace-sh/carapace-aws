package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_createDataSourceCmd = &cobra.Command{
	Use:   "create-data-source",
	Short: "Creates an Amazon DataZone data source.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_createDataSourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_createDataSourceCmd).Standalone()

		datazone_createDataSourceCmd.Flags().String("asset-forms-input", "", "The metadata forms that are to be attached to the assets that this data source works with.")
		datazone_createDataSourceCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_createDataSourceCmd.Flags().String("configuration", "", "Specifies the configuration of the data source.")
		datazone_createDataSourceCmd.Flags().String("connection-identifier", "", "The ID of the connection.")
		datazone_createDataSourceCmd.Flags().String("description", "", "The description of the data source.")
		datazone_createDataSourceCmd.Flags().String("domain-identifier", "", "The ID of the Amazon DataZone domain where the data source is created.")
		datazone_createDataSourceCmd.Flags().String("enable-setting", "", "Specifies whether the data source is enabled.")
		datazone_createDataSourceCmd.Flags().String("environment-identifier", "", "The unique identifier of the Amazon DataZone environment to which the data source publishes assets.")
		datazone_createDataSourceCmd.Flags().String("name", "", "The name of the data source.")
		datazone_createDataSourceCmd.Flags().Bool("no-publish-on-import", false, "Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.")
		datazone_createDataSourceCmd.Flags().String("project-identifier", "", "The identifier of the Amazon DataZone project in which you want to add this data source.")
		datazone_createDataSourceCmd.Flags().Bool("publish-on-import", false, "Specifies whether the assets that this data source creates in the inventory are to be also automatically published to the catalog.")
		datazone_createDataSourceCmd.Flags().String("recommendation", "", "Specifies whether the business name generation is to be enabled for this data source.")
		datazone_createDataSourceCmd.Flags().String("schedule", "", "The schedule of the data source runs.")
		datazone_createDataSourceCmd.Flags().String("type", "", "The type of the data source.")
		datazone_createDataSourceCmd.MarkFlagRequired("domain-identifier")
		datazone_createDataSourceCmd.MarkFlagRequired("name")
		datazone_createDataSourceCmd.Flag("no-publish-on-import").Hidden = true
		datazone_createDataSourceCmd.MarkFlagRequired("project-identifier")
		datazone_createDataSourceCmd.MarkFlagRequired("type")
	})
	datazoneCmd.AddCommand(datazone_createDataSourceCmd)
}
