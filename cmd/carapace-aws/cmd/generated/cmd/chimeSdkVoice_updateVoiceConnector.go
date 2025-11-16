package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_updateVoiceConnectorCmd = &cobra.Command{
	Use:   "update-voice-connector",
	Short: "Updates the details for the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_updateVoiceConnectorCmd).Standalone()

	chimeSdkVoice_updateVoiceConnectorCmd.Flags().String("name", "", "The name of the Voice Connector.")
	chimeSdkVoice_updateVoiceConnectorCmd.Flags().Bool("no-require-encryption", false, "When enabled, requires encryption for the Voice Connector.")
	chimeSdkVoice_updateVoiceConnectorCmd.Flags().Bool("require-encryption", false, "When enabled, requires encryption for the Voice Connector.")
	chimeSdkVoice_updateVoiceConnectorCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_updateVoiceConnectorCmd.MarkFlagRequired("name")
	chimeSdkVoice_updateVoiceConnectorCmd.Flag("no-require-encryption").Hidden = true
	chimeSdkVoice_updateVoiceConnectorCmd.MarkFlagRequired("no-require-encryption")
	chimeSdkVoice_updateVoiceConnectorCmd.MarkFlagRequired("require-encryption")
	chimeSdkVoice_updateVoiceConnectorCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_updateVoiceConnectorCmd)
}
