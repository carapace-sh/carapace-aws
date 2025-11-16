package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backup_describeRegionSettingsCmd = &cobra.Command{
	Use:   "describe-region-settings",
	Short: "Returns the current service opt-in settings for the Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(backup_describeRegionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(backup_describeRegionSettingsCmd).Standalone()

	})
	backupCmd.AddCommand(backup_describeRegionSettingsCmd)
}
