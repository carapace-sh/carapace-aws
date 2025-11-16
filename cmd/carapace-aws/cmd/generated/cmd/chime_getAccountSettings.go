package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Retrieves account settings for the specified Amazon Chime account ID, such as remote control and dialout settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getAccountSettingsCmd).Standalone()

	chime_getAccountSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
	chime_getAccountSettingsCmd.MarkFlagRequired("account-id")
	chimeCmd.AddCommand(chime_getAccountSettingsCmd)
}
