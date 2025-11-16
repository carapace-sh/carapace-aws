package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateProxySessionCmd = &cobra.Command{
	Use:   "update-proxy-session",
	Short: "Updates the specified proxy session details, such as voice or SMS capabilities.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateProxySessionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_updateProxySessionCmd).Standalone()

		chimeSdkVoice_updateProxySessionCmd.Flags().String("capabilities", "", "The proxy session capabilities.")
		chimeSdkVoice_updateProxySessionCmd.Flags().String("expiry-minutes", "", "The number of minutes allowed for the proxy session.")
		chimeSdkVoice_updateProxySessionCmd.Flags().String("proxy-session-id", "", "The proxy session ID.")
		chimeSdkVoice_updateProxySessionCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_updateProxySessionCmd.MarkFlagRequired("capabilities")
		chimeSdkVoice_updateProxySessionCmd.MarkFlagRequired("proxy-session-id")
		chimeSdkVoice_updateProxySessionCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateProxySessionCmd)
}
