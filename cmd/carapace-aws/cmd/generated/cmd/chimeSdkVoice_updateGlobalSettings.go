package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateGlobalSettingsCmd = &cobra.Command{
	Use:   "update-global-settings",
	Short: "Updates global settings for the Amazon Chime SDK Voice Connectors in an AWS account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateGlobalSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_updateGlobalSettingsCmd).Standalone()

		chimeSdkVoice_updateGlobalSettingsCmd.Flags().String("voice-connector", "", "The Voice Connector settings.")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateGlobalSettingsCmd)
}
