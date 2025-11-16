package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteProxySessionCmd = &cobra.Command{
	Use:   "delete-proxy-session",
	Short: "Deletes the specified proxy session from the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteProxySessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteProxySessionCmd).Standalone()

		chimeSdkVoice_deleteProxySessionCmd.Flags().String("proxy-session-id", "", "The proxy session ID.")
		chimeSdkVoice_deleteProxySessionCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_deleteProxySessionCmd.MarkFlagRequired("proxy-session-id")
		chimeSdkVoice_deleteProxySessionCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteProxySessionCmd)
}
