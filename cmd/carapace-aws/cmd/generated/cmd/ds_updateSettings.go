package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ds_updateSettingsCmd = &cobra.Command{
	Use:   "update-settings",
	Short: "Updates the configurable settings for the specified directory.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ds_updateSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ds_updateSettingsCmd).Standalone()

		ds_updateSettingsCmd.Flags().String("directory-id", "", "The identifier of the directory for which to update settings.")
		ds_updateSettingsCmd.Flags().String("settings", "", "The list of [Setting]() objects.")
		ds_updateSettingsCmd.MarkFlagRequired("directory-id")
		ds_updateSettingsCmd.MarkFlagRequired("settings")
	})
	dsCmd.AddCommand(ds_updateSettingsCmd)
}
