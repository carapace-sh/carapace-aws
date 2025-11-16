package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getVoiceConnectorProxyCmd = &cobra.Command{
	Use:   "get-voice-connector-proxy",
	Short: "Retrieves the proxy configuration details for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getVoiceConnectorProxyCmd).Standalone()

	chimeSdkVoice_getVoiceConnectorProxyCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getVoiceConnectorProxyCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getVoiceConnectorProxyCmd)
}
