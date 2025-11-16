package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var chimeSdkVoice_createSipMediaApplicationCallCmd = &cobra.Command{
	Use:   "create-sip-media-application-call",
	Short: "Creates an outbound call to a phone number from the phone number specified in the request, and it invokes the endpoint of the specified `sipMediaApplicationId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(chimeSdkVoice_createSipMediaApplicationCallCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(chimeSdkVoice_createSipMediaApplicationCallCmd).Standalone()

		chimeSdkVoice_createSipMediaApplicationCallCmd.Flags().String("arguments-map", "", "Context passed to a CreateSipMediaApplication API call.")
		chimeSdkVoice_createSipMediaApplicationCallCmd.Flags().String("from-phone-number", "", "The phone number that a user calls from.")
		chimeSdkVoice_createSipMediaApplicationCallCmd.Flags().String("sip-headers", "", "The SIP headers added to an outbound call leg.")
		chimeSdkVoice_createSipMediaApplicationCallCmd.Flags().String("sip-media-application-id", "", "The ID of the SIP media application.")
		chimeSdkVoice_createSipMediaApplicationCallCmd.Flags().String("to-phone-number", "", "The phone number that the service should call.")
		chimeSdkVoice_createSipMediaApplicationCallCmd.MarkFlagRequired("from-phone-number")
		chimeSdkVoice_createSipMediaApplicationCallCmd.MarkFlagRequired("sip-media-application-id")
		chimeSdkVoice_createSipMediaApplicationCallCmd.MarkFlagRequired("to-phone-number")
	})
	chimeSdkVoiceCmd.AddCommand(chimeSdkVoice_createSipMediaApplicationCallCmd)
}
