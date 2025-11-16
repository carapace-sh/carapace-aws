package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_putVoiceConnectorOriginationCmd = &cobra.Command{
	Use:   "put-voice-connector-origination",
	Short: "Updates a Voice Connector's origination settings.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_putVoiceConnectorOriginationCmd).Standalone()

	chimeSdkVoice_putVoiceConnectorOriginationCmd.Flags().String("origination", "", "The origination settings being updated.")
	chimeSdkVoice_putVoiceConnectorOriginationCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_putVoiceConnectorOriginationCmd.MarkFlagRequired("origination")
	chimeSdkVoice_putVoiceConnectorOriginationCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_putVoiceConnectorOriginationCmd)
}
