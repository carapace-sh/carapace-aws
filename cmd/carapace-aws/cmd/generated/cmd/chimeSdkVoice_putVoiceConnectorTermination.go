package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorTerminationCmd = &cobra.Command{
	Use:   "put-voice-connector-termination",
	Short: "Updates a Voice Connector's termination settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorTerminationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_putVoiceConnectorTerminationCmd).Standalone()

		chimeSdkVoice_putVoiceConnectorTerminationCmd.Flags().String("termination", "", "The termination settings to be updated.")
		chimeSdkVoice_putVoiceConnectorTerminationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
		chimeSdkVoice_putVoiceConnectorTerminationCmd.MarkFlagRequired("termination")
		chimeSdkVoice_putVoiceConnectorTerminationCmd.MarkFlagRequired("voice-connector-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorTerminationCmd)
}
