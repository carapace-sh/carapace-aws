package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd = &cobra.Command{
	Use:   "delete-voice-connector-termination-credentials",
	Short: "Deletes the specified SIP credentials used by your equipment to authenticate during call termination.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd).Standalone()

		chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd.Flags().String("usernames", "", "The RFC2617 compliant username associated with the SIP credentials, in US-ASCII format.")
		chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd.MarkFlagRequired("usernames")
		chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_deleteVoiceConnectorTerminationCredentialsCmd)
}
