package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "get-voice-connector-group",
	Short: "Retrieves details for the specified Amazon Chime SDK Voice Connector group, such as timestamps,name, and associated `VoiceConnectorItems`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorGroupCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorGroupCmd.Flags().String("voice-connector-group-id", "", "The Voice Connector group ID.")
	chimeSdkVoice_getVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-group-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorGroupCmd)
}
