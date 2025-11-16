package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "disassociate-phone-numbers-from-voice-connector-group",
	Short: "Disassociates the specified phone numbers from the specified Amazon Chime SDK Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd).Standalone()

	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd.Flags().String("e164-phone-numbers", "", "The list of phone numbers, in E.164 format.")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd.Flags().String("voice-connector-group-id", "", "The Voice Connector group ID.")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd.MarkFlagRequired("e164-phone-numbers")
	chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-group-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_disassociatePhoneNumbersFromVoiceConnectorGroupCmd)
}
