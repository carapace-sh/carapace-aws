package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var artifact_getAccountSettingsCmd = &cobra.Command{
	Use:   "get-account-settings",
	Short: "Get the account settings for Artifact.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_getAccountSettingsCmd).Standalone()

	artifactCmd.AddCommand(artifact_getAccountSettingsCmd)
}
