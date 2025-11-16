package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorProxyCmd = &cobra.Command{
	Use:   "delete-voice-connector-proxy",
	Short: "Deletes the proxy configuration from the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorProxyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteVoiceConnectorProxyCmd).Standalone()

		chimeSdkVoice_deleteVoiceConnectorProxyCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_deleteVoiceConnectorProxyCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorProxyCmd)
}
