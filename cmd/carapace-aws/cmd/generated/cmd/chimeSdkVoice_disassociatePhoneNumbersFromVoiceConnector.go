package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd = &cobra.Command{
	Use:   "disassociate-phone-numbers-from-voice-connector",
	Short: "Disassociates the specified phone numbers from the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd).Standalone()

	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd.Flags().String("e164-phone-numbers", "", "List of phone numbers, in E.164 format.")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd.MarkFlagRequired("e164-phone-numbers")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorCmd)
}
