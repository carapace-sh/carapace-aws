package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd = &cobra.Command{
	Use:   "put-voice-connector-termination-credentials",
	Short: "Updates a Voice Connector's termination credentials.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd).Standalone()

		chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd.Flags().String("credentials", "", "The termination credentials being updated.")
		chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorTerminationCredentialsCmd)
}
