package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_listVoiceConnectorTerminationCredentialsCmd = &cobra.Command{
	Use:   "list-voice-connector-termination-credentials",
	Short: "Lists the SIP credentials for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_listVoiceConnectorTerminationCredentialsCmd).Standalone()

	chimeSdkVoice_listVoiceConnectorTerminationCredentialsCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_listVoiceConnectorTerminationCredentialsCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_listVoiceConnectorTerminationCredentialsCmd)
}
