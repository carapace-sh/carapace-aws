package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd = &cobra.Command{
	Use:   "associate-phone-numbers-with-voice-connector-group",
	Short: "Associates phone numbers with the specified Amazon Chime SDK Voice Connector group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd).Standalone()

		chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd.Flags().String("e164-phone-numbers", "", "List of phone numbers, in E.164 format.")
		chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd.Flags().String("force-associate", "", "If true, associates the provided phone numbers with the provided Amazon Chime SDK Voice Connector Group and removes any previously existing associations.")
		chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd.Flags().String("voice-connector-group-id", "", "The Amazon Chime SDK Voice Connector group ID.")
		chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd.MarkFlagRequired("e164-phone-numbers")
		chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd.MarkFlagRequired("voice-connector-group-id")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_associatePhoneNumbersWithVoiceConnectorGroupCmd)
}
