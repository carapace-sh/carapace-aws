package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Returns the number of unmetered iOS or unmetered Android devices that have been purchased by the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_getAccountSettingsCmd).Standalone()

	devicefarmCmd.AddCommand(devicefarm_getAccountSettingsCmd)
}
