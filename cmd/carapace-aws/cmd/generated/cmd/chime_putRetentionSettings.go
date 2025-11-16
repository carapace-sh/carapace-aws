package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_putRetentionSettingsCmd = &cobra.Command{
	Use:   "put-retention-settings",
	Short: "Puts retention settings for the specified Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_putRetentionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_putRetentionSettingsCmd).Standalone()

		chime_putRetentionSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_putRetentionSettingsCmd.Flags().String("retention-settings", "", "The retention settings.")
		chime_putRetentionSettingsCmd.MarkFlagRequired("account-id")
		chime_putRetentionSettingsCmd.MarkFlagRequired("retention-settings")
	})
	chimeCmd.AddCommand(chime_putRetentionSettingsCmd)
}
