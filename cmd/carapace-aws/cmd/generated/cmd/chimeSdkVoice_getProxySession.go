package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_getProxySessionCmd = &cobra.Command{
	Use:   "get-proxy-session",
	Short: "Retrieves the specified proxy session details for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_getProxySessionCmd).Standalone()

	chimeSdkVoice_getProxySessionCmd.Flags().String("proxy-session-id", "", "The proxy session ID.")
	chimeSdkVoice_getProxySessionCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_getProxySessionCmd.MarkFlagRequired("proxy-session-id")
	chimeSdkVoice_getProxySessionCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_getProxySessionCmd)
}
