package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_updateUserSettingsCmd = &cobra.Command{
	Use:   "update-user-settings",
	Short: "Updates the settings for the specified user, such as phone number settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_updateUserSettingsCmd).Standalone()

	chime_updateUserSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_updateUserSettingsCmd.Flags().String("user-id", "", "The user ID.")
	chime_updateUserSettingsCmd.Flags().String("user-settings", "", "The user settings to update.")
	chime_updateUserSettingsCmd.MarkFlagRequired("account-id")
	chime_updateUserSettingsCmd.MarkFlagRequired("user-id")
	chime_updateUserSettingsCmd.MarkFlagRequired("user-settings")
	chimeCmd.AddCommand(chime_updateUserSettingsCmd)
}
