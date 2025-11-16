package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorCmd = &cobra.Command{
	Use:   "get-voice-connector",
	Short: "Retrieves details for the specified Amazon Chime SDK Voice Connector, such as timestamps,name, outbound host, and encryption requirements.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorCmd)
}
