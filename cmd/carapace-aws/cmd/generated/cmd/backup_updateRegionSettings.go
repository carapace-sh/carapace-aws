package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateRegionSettingsCmd = &cobra.Command{
	Use:   "update-region-settings",
	Short: "Updates the current service opt-in settings for the Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateRegionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_updateRegionSettingsCmd).Standalone()

		backup_updateRegionSettingsCmd.Flags().String("resource-type-management-preference", "", "Enables or disables full Backup management of backups for a resource type.")
		backup_updateRegionSettingsCmd.Flags().String("resource-type-opt-in-preference", "", "Updates the list of services along with the opt-in preferences for the Region.")
	})
	backupCmd.AddCommand(backup_updateRegionSettingsCmd)
}
