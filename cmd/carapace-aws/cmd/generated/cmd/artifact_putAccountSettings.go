package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_putAccountSettingsCmd = &cobra.Command{
	Use:   "put-account-settings",
	Short: "Put the account settings for Artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_putAccountSettingsCmd).Standalone()

	artifact_putAccountSettingsCmd.Flags().String("notification-subscription-status", "", "Desired notification subscription status.")
	artifactCmd.AddCommand(artifact_putAccountSettingsCmd)
}
