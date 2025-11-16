package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_updateGlobalSettingsCmd = &cobra.Command{
	Use:   "update-global-settings",
	Short: "Updates whether the Amazon Web Services account is opted in to cross-account backup.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_updateGlobalSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_updateGlobalSettingsCmd).Standalone()

		backup_updateGlobalSettingsCmd.Flags().String("global-settings", "", "Inputs can include:")
	})
	backupCmd.AddCommand(backup_updateGlobalSettingsCmd)
}
