package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chime_getRetentionSettingsCmd = &cobra.Command{
	Use:   "get-retention-settings",
	Short: "Gets the retention settings for the specified Amazon Chime Enterprise account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chime_getRetentionSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chime_getRetentionSettingsCmd).Standalone()

		chime_getRetentionSettingsCmd.Flags().String("account-id", "", "The Amazon Chime account ID.")
		chime_getRetentionSettingsCmd.MarkFlagRequired("account-id")
	})
	chimeCmd.AddCommand(chime_getRetentionSettingsCmd)
}
