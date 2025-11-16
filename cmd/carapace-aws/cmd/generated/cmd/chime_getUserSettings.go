package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getUserSettingsCmd = &cobra.Command{
	Use:   "get-user-settings",
	Short: "Retrieves settings for the specified user ID, such as any associated phone number settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getUserSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_getUserSettingsCmd).Standalone()

		chime_getUserSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_getUserSettingsCmd.Flags().String("user-id", "", "The user ID.")
		chime_getUserSettingsCmd.MarkFlagRequired("account-id")
		chime_getUserSettingsCmd.MarkFlagRequired("user-id")
	})
	chimeCmd.AddCommand(chime_getUserSettingsCmd)
}
