package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_createDataProviderCmd = &cobra.Command{
	Use:   "create-data-provider",
	Short: "Creates a data provider using the provided settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_createDataProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_createDataProviderCmd).Standalone()

		dms_createDataProviderCmd.Flags().String("data-provider-name", "", "A user-friendly name for the data provider.")
		dms_createDataProviderCmd.Flags().String("description", "", "A user-friendly description of the data provider.")
		dms_createDataProviderCmd.Flags().String("engine", "", "The type of database engine for the data provider.")
		dms_createDataProviderCmd.Flags().String("settings", "", "The settings in JSON format for a data provider.")
		dms_createDataProviderCmd.Flags().String("tags", "", "One or more tags to be assigned to the data provider.")
		dms_createDataProviderCmd.Flags().String("virtual", "", "Indicates whether the data provider is virtual.")
		dms_createDataProviderCmd.MarkFlagRequired("engine")
		dms_createDataProviderCmd.MarkFlagRequired("settings")
	})
	dmsCmd.AddCommand(dms_createDataProviderCmd)
}
