package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getGlobalSettingsCmd = &cobra.Command{
	Use:   "get-global-settings",
	Short: "Retrieves the global settings for the Amazon Chime SDK Voice Connectors in an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getGlobalSettingsCmd).Standalone()

	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getGlobalSettingsCmd)
}
