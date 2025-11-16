package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateAccountSettingsCmd = &cobra.Command{
	Use:   "update-account-settings",
	Short: "Updates the settings for the specified Amazon Chime account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateAccountSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_updateAccountSettingsCmd).Standalone()

		chime_updateAccountSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_updateAccountSettingsCmd.Flags().String("account-settings", "", "The Amazon Chime account settings to update.")
		chime_updateAccountSettingsCmd.MarkFlagRequired("account-id")
		chime_updateAccountSettingsCmd.MarkFlagRequired("account-settings")
	})
	chimeCmd.AddCommand(chime_updateAccountSettingsCmd)
}
