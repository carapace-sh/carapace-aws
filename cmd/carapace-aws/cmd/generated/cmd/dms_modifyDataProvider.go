package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dms_modifyDataProviderCmd = &cobra.Command{
	Use:   "modify-data-provider",
	Short: "Modifies the specified data provider using the provided settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dms_modifyDataProviderCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dms_modifyDataProviderCmd).Standalone()

		dms_modifyDataProviderCmd.Flags().String("data-provider-identifier", "", "The identifier of the data provider.")
		dms_modifyDataProviderCmd.Flags().String("data-provider-name", "", "The name of the data provider.")
		dms_modifyDataProviderCmd.Flags().String("description", "", "A user-friendly description of the data provider.")
		dms_modifyDataProviderCmd.Flags().String("engine", "", "The type of database engine for the data provider.")
		dms_modifyDataProviderCmd.Flags().String("exact-settings", "", "If this attribute is Y, the current call to `ModifyDataProvider` replaces all existing data provider settings with the exact settings that you specify in this call.")
		dms_modifyDataProviderCmd.Flags().String("settings", "", "The settings in JSON format for a data provider.")
		dms_modifyDataProviderCmd.Flags().String("virtual", "", "Indicates whether the data provider is virtual.")
		dms_modifyDataProviderCmd.MarkFlagRequired("data-provider-identifier")
	})
	dmsCmd.AddCommand(dms_modifyDataProviderCmd)
}
