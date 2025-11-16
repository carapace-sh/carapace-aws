package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createVoiceConnectorCmd = &cobra.Command{
	Use:   "create-voice-connector",
	Short: "Creates an Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createVoiceConnectorCmd).Standalone()

	chimeSdkVoice_createVoiceConnectorCmd.Flags().String("aws-region", "", "The AWS Region in which the Amazon Chime SDK Voice Connector is created.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().String("integration-type", "", "The connectors for use with Amazon Connect.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().String("name", "", "The name of the Voice Connector.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().String("network-type", "", "The type of network for the Voice Connector.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().Bool("no-require-encryption", false, "Enables or disables encryption for the Voice Connector.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().Bool("require-encryption", false, "Enables or disables encryption for the Voice Connector.")
	chimeSdkVoice_createVoiceConnectorCmd.Flags().String("tags", "", "The tags assigned to the Voice Connector.")
	chimeSdkVoice_createVoiceConnectorCmd.MarkFlagRequired("name")
	chimeSdkVoice_createVoiceConnectorCmd.Flag("no-require-encryption").Hidden = true
	chimeSdkVoice_createVoiceConnectorCmd.MarkFlagRequired("no-require-encryption")
	chimeSdkVoice_createVoiceConnectorCmd.MarkFlagRequired("require-encryption")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createVoiceConnectorCmd)
}
