package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd = &cobra.Command{
	Use:   "associate-phone-numbers-with-voice-connector",
	Short: "Associates phone numbers with the specified Amazon Chime SDK Voice Connector.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd).Standalone()

	chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd.Flags().String("e164-phone-numbers", "", "List of phone numbers, in E.164 format.")
	chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd.Flags().String("force-associate", "", "If true, associates the provided phone numbers with the provided Amazon Chime SDK Voice Connector and removes any previously existing associations.")
	chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd.Flags().String("voice-connector-id", "", "The Voice Connector ID.")
	chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd.MarkFlagRequired("e164-phone-numbers")
	chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd.MarkFlagRequired("voice-connector-id")
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorCmd)
}
