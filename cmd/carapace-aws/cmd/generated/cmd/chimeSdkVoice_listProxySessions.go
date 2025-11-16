package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listProxySessionsCmd = &cobra.Command{
	Use:   "list-proxy-sessions",
	Short: "Lists the proxy sessions for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listProxySessionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_listProxySessionsCmd).Standalone()

		chimeSdkVoice_listProxySessionsCmd.Flags().String("max-results", "", "The maximum number of results to return in a single call.")
		chimeSdkVoice_listProxySessionsCmd.Flags().String("next-token", "", "The token used to retrieve the next page of results.")
		chimeSdkVoice_listProxySessionsCmd.Flags().String("status", "", "The proxy session status.")
		chimeSdkVoice_listProxySessionsCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_listProxySessionsCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listProxySessionsCmd)
}
